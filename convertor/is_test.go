package convertor

import (
	"convert/apistructures"
	"convert/targetstructures"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_IsNew(t *testing.T) {
	testItems := apistructures.Samples()["multi"]
	items := map[string]targetstructures.Item{}

	for _, i := range testItems {
		tran := Transform(apistructures.DataFixes(i))
		k := tran.Attributes.URIS.Slug
		items[k] = tran
	}

	// abalone has seasonable times
	abalone := items["abalone"]
	// north is new, south isnt
	june := time.Date(2021, time.Month(6), 1, 1, 0, 0, 0, time.UTC)
	is := GenerateIs(june, abalone)
	assert.True(t, is["northern"]["new"])
	assert.False(t, is["southern"]["new"])

	// north isnt new, south is
	july := time.Date(2021, time.Month(7), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(july, abalone)
	assert.False(t, is["northern"]["new"])
	assert.False(t, is["southern"]["new"])

	feb := time.Date(2021, time.Month(2), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(feb, abalone)
	assert.False(t, is["northern"]["new"])
	assert.False(t, is["southern"]["new"])

	dec := time.Date(2021, time.Month(12), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(dec, abalone)
	assert.False(t, is["northern"]["new"])
	assert.True(t, is["southern"]["new"])

	// base is available all year, so should never be true
	bass := items["sea-bass"]
	for i := 1; i <= 12; i++ {
		m := time.Date(2021, time.Month(i), 1, 1, 0, 0, 0, time.UTC)
		is = GenerateIs(m, bass)
		assert.False(t, is["northern"]["new"])
		assert.False(t, is["southern"]["new"])
	}
}

func Test_Leaving(t *testing.T) {
	testItems := apistructures.Samples()["multi"]
	items := map[string]targetstructures.Item{}

	for _, i := range testItems {
		tran := Transform(apistructures.DataFixes(i))
		k := tran.Attributes.URIS.Slug
		items[k] = tran
	}

	// abalone has seasons
	abalone := items["abalone"]
	jan := time.Date(2021, time.Month(1), 1, 1, 0, 0, 0, time.UTC)
	is := GenerateIs(jan, abalone)
	assert.True(t, is["northern"]["leaving"])
	assert.False(t, is["southern"]["leaving"])

	july := time.Date(2021, time.Month(7), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(july, abalone)
	assert.False(t, is["northern"]["leaving"])
	assert.True(t, is["southern"]["leaving"])

	march := time.Date(2021, time.Month(3), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(march, abalone)
	assert.False(t, is["northern"]["leaving"])
	assert.False(t, is["southern"]["leaving"])

	// base is available all year, so should never be true
	bass := items["sea-bass"]
	for i := 1; i <= 12; i++ {
		m := time.Date(2021, time.Month(i), 1, 1, 0, 0, 0, time.UTC)
		is = GenerateIs(m, bass)
		assert.False(t, is["northern"]["leaving"])
		assert.False(t, is["southern"]["leaving"])
	}
}

func Test_Available(t *testing.T) {
	testItems := apistructures.Samples()["multi"]
	items := map[string]targetstructures.Item{}

	for _, i := range testItems {
		tran := Transform(apistructures.DataFixes(i))
		k := tran.Attributes.URIS.Slug
		items[k] = tran
	}

	abalone := items["abalone"]
	jan := time.Date(2021, time.Month(1), 1, 1, 0, 0, 0, time.UTC)
	is := GenerateIs(jan, abalone)
	assert.True(t, is["northern"]["available"])
	assert.True(t, is["southern"]["available"])

	march := time.Date(2021, time.Month(3), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(march, abalone)

	assert.False(t, is["northern"]["available"])
	assert.True(t, is["southern"]["available"])

	// base is always available
	bass := items["sea-bass"]
	for i := 1; i <= 12; i++ {
		m := time.Date(2021, time.Month(i), 1, 1, 0, 0, 0, time.UTC)
		is = GenerateIs(m, bass)
		assert.True(t, is["northern"]["available"])
		assert.True(t, is["southern"]["available"])
	}

}
