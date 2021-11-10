package helpers

import (
	"sort"
	"sync"
)

type TopNEntry struct {
	Id      string
	Counter int
}
type TopNEntries []TopNEntry
type topNIndex map[string]int

type TopN struct {
	index      topNIndex
	entries    TopNEntries
	accessLock sync.Mutex
}

func NewTopN() *TopN {
	var t TopN
	t.Clear()
	return &t
}

func (t *TopN) Clear() {
	t.index = make(topNIndex)
	t.entries = make(TopNEntries, 100) // TODO: meaningful initial size
}

func (t *TopN) Add(id string, weight int) {
	// probably overkill but debugging it later is hard...
	t.accessLock.Lock()
	defer t.accessLock.Unlock()

	pos, idExists := t.index[id]
	if !idExists {
		entry := TopNEntry{
			Id:      id,
			Counter: weight,
		}
		t.index[id] = len(t.entries)
		t.entries = append(t.entries, entry)
	} else {
		t.entries[pos].Counter = t.entries[pos].Counter + weight
	}
}

// sorting the slice destroys the index
// so either clear or rebuild the index, only implemented the first
func (t *TopN) GetTopNAndClear(n int) TopNEntries {
	t.accessLock.Lock()
	defer t.Clear()
	defer t.accessLock.Unlock()

	sort.Slice(t.entries, func(i, j int) bool {
		return t.entries[i].Counter > t.entries[j].Counter
	})

	return t.entries[:n]
}
