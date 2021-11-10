package helpers

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type CommitSizer struct {
	entries map[string]int
}

func LoadCommitSizer(filename string) *CommitSizer {
	var l CommitSizer
	l.entries = make(map[string]int)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	for {

		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// sha,message,event_id
		event_id := record[2]

		l.entries[event_id] = l.entries[event_id] + 1
	}
	return &l
}

func (l CommitSizer) Length() int {
	return len(l.entries)
}

func (l CommitSizer) GetCommitSize(event_id string) int {
	commit_size, _ := l.entries[event_id]
	return commit_size
}
