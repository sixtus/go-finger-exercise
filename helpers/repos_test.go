package helpers

import (
	"testing"
)

func Test_Repos_Lookup(t *testing.T) {
	reposFileName := "../test-data/repos.csv"

	a := LoadLookup(&reposFileName)

	expectedLength := 12607 // this is specific to the test data

	testId := "224252202" // again specific to test data
	testExpected := "DSC-RPI/dsc-portal"

	if a.Length() != expectedLength {
		t.Fatalf("There are %d elements, but expected %d", a.Length(), expectedLength)
	}

	if a.GetNameById(testId) != testExpected {
		t.Fatalf("Looking up %s, expected %s, got %s", testId, testExpected, a.GetNameById(testId))
	}
}
