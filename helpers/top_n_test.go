package helpers

import (
	"testing"
)

func Test_TopN(t *testing.T) {
	top := NewTopN()

	top.Add("A", 5)
	top.Add("A", 1)
	top.Add("B", 2)
	top.Add("B", 1)
	top.Add("C", 1)
	top.Add("C", 10)

	expectedLength := 2
	testResponse := top.GetTopNAndClear(expectedLength)

	if len(testResponse) != expectedLength {
		t.Fatalf("There are %d elements, but expected %d", len(testResponse), expectedLength)
	}

	if testResponse[0].Name != "C" || testResponse[0].Counter != 11 {
		t.Fatalf("1st element doesn't match, expected C 11, got %s %d", testResponse[0].Name, testResponse[0].Counter)
	}

	if testResponse[1].Name != "A" || testResponse[1].Counter != 6 {
		t.Fatalf("2nd element doesn't match, expected A 6, got %s %d", testResponse[1].Name, testResponse[1].Counter)
	}

	testResponse = top.GetTopNAndClear(expectedLength)
	if testResponse[0].Name != "" || testResponse[0].Counter != 0 {
		t.Fatalf("GetTopNAndClear didn't clear, got %s %d as top", testResponse[0].Name, testResponse[0].Counter)
	}

}
