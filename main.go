package main

import (
	"flag"
	"fmt"
)

func main() {
	actorsFileName := flag.String("actors", "test-data/actors.csv", "the actors.csv file")
	commitsFileName := flag.String("commits", "test-data/commits.csv", "the commits.csv file")
	eventsFileName := flag.String("events", "test-data/events.csv", "the events.csv file")
	reposFileName := flag.String("repos", "test-data/repots.csv", "the repos.csv file")

	flag.Parse()

	fmt.Println(*actorsFileName)
	fmt.Println(*commitsFileName)
	fmt.Println(*eventsFileName)
	fmt.Println(*reposFileName)
}
