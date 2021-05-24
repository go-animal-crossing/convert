package apistructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DataFixes(t *testing.T) {
	test := Samples()["multi"]
	abalone := Item{}

	for _, i := range test {
		if i.FileName == "abalone" {
			abalone = i
		}
	}

	assert.Equal(t, "abalone", abalone.FileName)

	fixed := DataFixes(abalone)
	expected := []int{12, 1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, fixed.Availability.MonthArraySouthern)
}
