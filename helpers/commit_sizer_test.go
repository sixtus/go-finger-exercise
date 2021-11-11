package helpers

import (
	"testing"
)

func Test_Commit_Sizer(t *testing.T) {
	commitsFileName := "../test-data/commits.csv"

	c := LoadCommitSizer(commitsFileName)

	expectedLength := 16520 // this is specific to the test data

	if c.Length() != expectedLength {
		t.Fatalf("There are %d elements, but expected %d", c.Length(), expectedLength)
	}

	testEvent := "11185376329"
	expectedSize := 4

	if c.GetCommitSize(testEvent) != expectedSize {
		t.Fatalf("For got %d events for id %s, expected %d", c.GetCommitSize(testEvent), testEvent, expectedSize)
	}

}
