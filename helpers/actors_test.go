package helpers

import (
	"testing"
)

func Test_Actor(t *testing.T) {
	actorsFileName := "../test-data/actors.csv"

	a := LoadActors(actorsFileName)

	expectedLength := 9728 // this is specific to the test data

	testKey := "8422699" // again specific to test data
	testExpected := "Apexal"

	if a.Length() != expectedLength {
		t.Fatalf("There are %d elements, but expected %d", a.Length(), expectedLength)
	}

	if a.Lookup(testKey) != testExpected {
		t.Fatalf("Looking up %s, expected %s, got %s", testKey, testExpected, a.Lookup(testKey))
	}
}
