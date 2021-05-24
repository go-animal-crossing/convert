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
		tran := transform(apistructures.DataFixes(i))
		k := tran.Attributes.URIS.Slug
		items[k] = tran
	}

	// abalone has seasonable times
	abalone := items["abalone"]
	// north is new, south isnt
	june := time.Date(2021, time.Month(6), 1, 1, 0, 0, 0, time.UTC)
	is := GenerateIs(june, abalone)
	assert.True(t, is.Northern.New)
	assert.False(t, is.Southern.New)

	// north isnt new, south is
	july := time.Date(2021, time.Month(7), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(july, abalone)
	assert.False(t, is.Northern.New)
	assert.False(t, is.Southern.New)

	feb := time.Date(2021, time.Month(2), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(feb, abalone)
	assert.False(t, is.Northern.New)
	assert.False(t, is.Southern.New)

	dec := time.Date(2021, time.Month(12), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(dec, abalone)
	assert.False(t, is.Northern.New)
	assert.True(t, is.Southern.New)

	// base is available all year, so should never be true
	bass := items["sea-bass"]
	for i := 1; i <= 12; i++ {
		m := time.Date(2021, time.Month(i), 1, 1, 0, 0, 0, time.UTC)
		is = GenerateIs(m, bass)
		assert.False(t, is.Northern.New)
		assert.False(t, is.Southern.New)
	}
}

func Test_Leaving(t *testing.T) {
	testItems := apistructures.Samples()["multi"]
	items := map[string]targetstructures.Item{}

	for _, i := range testItems {
		tran := transform(apistructures.DataFixes(i))
		k := tran.Attributes.URIS.Slug
		items[k] = tran
	}

	// abalone has seasons
	abalone := items["abalone"]
	jan := time.Date(2021, time.Month(1), 1, 1, 0, 0, 0, time.UTC)
	is := GenerateIs(jan, abalone)
	assert.True(t, is.Northern.Leaving)
	assert.False(t, is.Southern.Leaving)

	july := time.Date(2021, time.Month(7), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(july, abalone)
	assert.False(t, is.Northern.Leaving)
	assert.True(t, is.Southern.Leaving)

	march := time.Date(2021, time.Month(3), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(march, abalone)
	assert.False(t, is.Northern.Leaving)
	assert.False(t, is.Southern.Leaving)

	// base is available all year, so should never be true
	bass := items["sea-bass"]
	for i := 1; i <= 12; i++ {
		m := time.Date(2021, time.Month(i), 1, 1, 0, 0, 0, time.UTC)
		is = GenerateIs(m, bass)
		assert.False(t, is.Northern.Leaving)
		assert.False(t, is.Southern.Leaving)
	}
}

func Test_Available(t *testing.T) {
	testItems := apistructures.Samples()["multi"]
	items := map[string]targetstructures.Item{}

	for _, i := range testItems {
		tran := transform(apistructures.DataFixes(i))
		k := tran.Attributes.URIS.Slug
		items[k] = tran
	}

	abalone := items["abalone"]
	jan := time.Date(2021, time.Month(1), 1, 1, 0, 0, 0, time.UTC)
	is := GenerateIs(jan, abalone)
	assert.True(t, is.Northern.Availabile)
	assert.True(t, is.Southern.Availabile)

	march := time.Date(2021, time.Month(3), 1, 1, 0, 0, 0, time.UTC)
	is = GenerateIs(march, abalone)

	assert.False(t, is.Northern.Availabile)
	assert.True(t, is.Southern.Availabile)

	// base is always available
	bass := items["sea-bass"]
	for i := 1; i <= 12; i++ {
		m := time.Date(2021, time.Month(i), 1, 1, 0, 0, 0, time.UTC)
		is = GenerateIs(m, bass)
		assert.True(t, is.Northern.Availabile)
		assert.True(t, is.Southern.Availabile)
	}

}
