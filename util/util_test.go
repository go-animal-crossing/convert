package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SafeQuote(t *testing.T) {
	original := `has "quotes" in 'it'`
	expected := "has quotes in it"
	actual := Safe(original)
	assert.Equal(t, expected, actual, "Should be equal")
}

func Test_Title(t *testing.T) {
	original := `is a title`
	expected := "Is A Title"
	actual := Title(original)
	assert.Equal(t, expected, actual, "Should be equal")
}

func Test_SlugSingle(t *testing.T) {
	original := `Is A Single`
	expected := "is-a-single"
	actual := Slugify(original)
	assert.Equal(t, expected, actual, "Should be equal")
}
func Test_URL(t *testing.T) {
	one := "Top level"
	two := "2nd level"
	expected := "/top-level/2nd-level"
	actual := URL(one, two)
	assert.Equal(t, expected, actual, "Should be equal")
}
