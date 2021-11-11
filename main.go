package main

import (
	"flag"
	"fmt"
	"github.com/sixtus/go-finger-exercise/internal"
)

const (
	N = 10
)

func output(label string, topN internal.TopNEntries, lookup *internal.Lookup) {
	fmt.Println(label)
	for n, e := range topN {
		fmt.Printf("%2d: %s -> %d\n", n+1, lookup.GetNameById(e.Id), e.Counter)
	}
	fmt.Println()
}

func main() {
	actorsFileName := flag.String("actors", "test-data/actors.csv", "the actors.csv file")
	commitsFileName := flag.String("commits", "test-data/commits.csv", "the commits.csv file")
	eventsFileName := flag.String("events", "test-data/events.csv", "the events.csv file")
	reposFileName := flag.String("repos", "test-data/repos.csv", "the repos.csv file")

	flag.Parse()

	actors := internal.LoadLookup(*actorsFileName)
	commits := internal.LoadCommitSizer(*commitsFileName)
	events := internal.LoadEventsScanner(*eventsFileName, commits)
	repos := internal.LoadLookup(*reposFileName)

	output("user by commits", events.TopUserCommits.GetTopNAndClear(N), actors)
	output("user by PRs", events.TopUserPR.GetTopNAndClear(N), actors)
	output("repo by commits", events.TopRepoCommits.GetTopNAndClear(N), repos)
	output("repo by watches", events.TopRepoWatches.GetTopNAndClear(N), repos)
}
