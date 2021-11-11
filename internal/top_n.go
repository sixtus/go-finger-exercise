package internal

import (
	"sort"
)

// The core of what we want to store, an id and it's counter
// TopN will sort based on the counter and it's fast to
// incrementally update the counter.
type TopNEntry struct {
	Id      string
	Counter int
}

// a slice of TopNEntry, the order in this slice is why we
// have this helper in the first place.
type TopNEntries []TopNEntry

type topNIndex map[string]int

// This is a helper for TopN queries ("what is the top 10 of....")
// The assumption is that you are streaming in ids via Add()
// Since we expect the adding in random (and repeated) order and
// access only once, the decision was taken against using "container/heap"
// (which does sorting on insertion and is not really designed for updates)
// Internally, TopN stores the entries in a slice, as this allows
// to use the built-in sorting mechanism of golang.
// The other design decision based on this assumption is to add an index,
// a map to improve the lookup in the entries (the map is not key->value,
// but key->entries-offset and the value is entries[entries-offset])
// This allows sorting to be delayed until the very end, while having good
// insertion speeds. The trade off is that once we sort, the index is invalid.
// For the given use case, it's perfectly enough to reset the content
// completely on reading, thus so far only a GetTopNAndClear() is exposed.
//
// This implenentation is not thread-safe.
//
// Note: a quickselect rather than a sort would be even faster, but
// the trade-off between last millisecond and drastically more complex
// code seems prudent here.
type TopN struct {
	index   topNIndex
	entries TopNEntries
}

func NewTopN() *TopN {
	var t TopN
	t.Clear()
	return &t
}

// re-initialize entries and index, also used by the constructor
func (t *TopN) Clear() {
	t.index = make(topNIndex)
	t.entries = make(TopNEntries, 0, 100) // TODO: meaningful initial size
}

// the primary input method. It's possible to give a weight, when
// in doubt, use 1.
func (t *TopN) Add(id string, weight int) {
	pos, idExists := t.index[id]
	if !idExists { // new id
		entry := TopNEntry{
			Id:      id,
			Counter: weight,
		}
		t.index[id] = len(t.entries)         // put the next slice position into the index
		t.entries = append(t.entries, entry) // append the new entry to the slice
	} else { // existing id
		t.entries[pos].Counter = t.entries[pos].Counter + weight
	}
}

// sorting the slice invalidates the index
// so either clear or rebuild the index, only implemented the first
func (t *TopN) GetTopNAndClear(n int) TopNEntries {
	defer t.Clear()

	sort.Slice(t.entries, func(i, j int) bool {
		return t.entries[i].Counter > t.entries[j].Counter
	})

	return t.entries[:n]
}
