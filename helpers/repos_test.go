package helpers

import (
	"testing"
)

func Test_Repos_Lookup(t *testing.T) {
	actorsFileName := "../test-data/repos.csv"

	a := LoadLookup(actorsFileName)

	expectedLength := 12607 // this is specific to the test data

	testKey := "224252202" // again specific to test data
	testExpected := "DSC-RPI/dsc-portal"

	if a.Length() != expectedLength {
		t.Fatalf("There are %d elements, but expected %d", a.Length(), expectedLength)
	}

	if a.GetName(testKey) != testExpected {
		t.Fatalf("Looking up %s, expected %s, got %s", testKey, testExpected, a.GetName(testKey))
	}
}