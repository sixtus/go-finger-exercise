package helpers

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Actors struct {
	lookup map[string]string
}

func LoadActors(filename string) *Actors {
	var a Actors
	a.lookup = make(map[string]string)

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

		actorId := record[0]
		userName := record[1]

		a.lookup[actorId] = userName
	}
	return &a
}

func (a Actors) Length() int {
	return len(a.lookup)
}

func (a Actors) Lookup(id string) string {
	name, _ := a.lookup[id]
	return name
}
