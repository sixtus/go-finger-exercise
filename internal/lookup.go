package internal

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Lookup struct {
	entries map[string]string
}

func LoadLookup(filename string) *Lookup {
	var l Lookup
	l.entries = make(map[string]string)

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

		id := record[0]
		humanReadable := record[1]

		l.entries[id] = humanReadable
	}
	return &l
}

func (l *Lookup) Length() int {
	return len(l.entries)
}

func (l *Lookup) GetNameById(id string) string {
	name, _ := l.entries[id]
	return name
}
