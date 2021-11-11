package internal

import (
	"testing"
)

const (
	N = 2
)

func check_reponse(t *testing.T, testCase string, is TopNEntries, expected TopNEntries) {
	if len(is) != len(expected) {
		t.Fatalf("Failed length in %s, is %d, expected %d", testCase, len(is), len(expected))
	}
	for i, e := range is {
		if e.Id != expected[i].Id {
			t.Fatalf("Failed id %d in %s, is %s, expected %s", i, testCase, e.Id, expected[i].Id)
		}
		if e.Counter != expected[i].Counter {
			t.Fatalf("Failed counter %d in %s, is %d, expected %d", i, testCase, e.Counter, expected[i].Counter)
		}
	}
}

func Test_EventsScanner(t *testing.T) {
	commits := LoadCommitSizer("../test-data/commits.csv")
	events := LoadEventsScanner("../test-data/events.csv", commits)

	check_reponse(t, "top user commits", events.TopUserCommits.GetTopNAndClear(N), TopNEntries{
		TopNEntry{
			Id:      "29139614",
			Counter: 451,
		},
		TopNEntry{
			Id:      "59293082",
			Counter: 331,
		},
	})

	check_reponse(t, "top user PRs", events.TopUserPR.GetTopNAndClear(N), TopNEntries{
		TopNEntry{
			Id:      "39814207",
			Counter: 256,
		},
		TopNEntry{
			Id:      "29139614",
			Counter: 245,
		},
	})

	check_reponse(t, "top repo commits", events.TopRepoCommits.GetTopNAndClear(N), TopNEntries{
		TopNEntry{
			Id:      "230501783",
			Counter: 331,
		},
		TopNEntry{
			Id:      "224857031",
			Counter: 222,
		},
	})

	check_reponse(t, "top repo watches", events.TopRepoWatches.GetTopNAndClear(N), TopNEntries{
		TopNEntry{
			Id:      "231135514",
			Counter: 44,
		},
		TopNEntry{
			Id:      "45307548",
			Counter: 11,
		},
	})
}
