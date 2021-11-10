package helpers

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type EventsScanner struct {
	TopUserPR      *TopN
	TopUserCommits *TopN
	TopRepoCommits *TopN
	TopRepoWatches *TopN
}

func LoadEventsScanner(filename *string, commits *CommitSizer) *EventsScanner {
	var e EventsScanner = EventsScanner{NewTopN(), NewTopN(), NewTopN(), NewTopN()}

	f, err := os.Open(*filename)
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

		// Note: I was torn on adjusting the naming convention to Go,
		//       opted for prefix "event_" + name in CVS.
		//       It 99% close to triggering creating a struct

		// CSV: id,type,actor_id,repo_
		event_id := record[0]
		event_type := record[1]
		event_actor_id := record[2]
		event_repo_id := record[3]

		switch event_type {
		case "PullRequestEvent":
			e.TopUserPR.Add(event_actor_id, 1)
		case "PushEvent":
			numberOfCommits := commits.GetCommitSize(event_id)
			e.TopUserCommits.Add(event_actor_id, numberOfCommits)
			e.TopRepoCommits.Add(event_repo_id, numberOfCommits)
		case "WatchEvent":
			e.TopRepoWatches.Add(event_repo_id, 1)
		}

	}
	return &e
}
