package maximumelement

import "testing"

func TestQueries(t *testing.T) {
	resultsArr := getMax(
		[]string{
			"1 97", "2", "1 20", "2", "1 26", "1 20", "2", "3", "1 91", "3",
		})

	if len(resultsArr) == 0 {
		t.Fail()
	}
}
