package convertor

import (
	"convert/apistructures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_titles(t *testing.T) {
	simple := apistructures.Samples()["simple"]
	actual := titles(simple)
	assert.Equal(t, simple.Names.EuEn, actual.Original, "Name should match")
	assert.Equal(t, "Bitterling", actual.Safe, "Title friendly name")
}

func Test_uris(t *testing.T) {
	simple := apistructures.Samples()["simple"]
	actual := uris(simple)
	assert.Equal(t, "bitterling", actual.Slug, "Slug should match")
	assert.Equal(t, "/fish/bitterling", actual.URL, "URL should be generated with type and slug")
}

func Test_prices(t *testing.T) {
	simple := apistructures.Samples()["simple"]
	actual := prices(simple)
	assert.Equal(t, 900, actual.Store, "Store should 900")
	assert.Equal(t, 1350, actual.Flick, "Flick should be 1350")
	assert.Equal(t, 0, actual.Cj, "CJ should be empty")
}

func Test_phrases(t *testing.T) {
	simple := apistructures.Samples()["simple"]
	actual := phrases(simple)
	assert.Equal(t, simple.CatchPhrase, actual.Capture.Original, "Capture catchpharse should match")
	assert.Equal(t, Safe(simple.CatchPhrase), actual.Capture.Safe, "Safe version should match safe result")
	assert.Equal(t, "Yes! Found it", actual.Capture.Safe, "Manual test for Safe")
}

func Test_images(t *testing.T) {
	simple := apistructures.Samples()["simple"]
	actual := images(simple)
	thumb := ImagePath(simple.Type, "thumb", simple.FileName, "png")
	assert.Equal(t, simple.IconURI, actual.Thumb.Direct, "Image paths should match")
	assert.Equal(t, thumb, actual.Thumb.Local, "Image paths should look like local")
	assert.Equal(t, "/fish/thumb/bitterling.png", actual.Thumb.Local, "Image paths should match what we think for local")
}

func Test_availability(t *testing.T) {
	simple := apistructures.Samples()["simple"]
	actual := availability(simple)

	assert.Equal(t, simple.Availability.Location, actual.Location, "Locations match")
	assert.Equal(t, simple.Availability.Rarity, actual.Rarity, "Rarity match")
	assert.False(t, actual.Months.Always, "Should not be available all year")

	// north
	assert.False(t, actual.Months.Northern.Always, "Should not be available all year in north")
	assert.Equal(t, simple.Availability.MonthArrayNorthern, actual.Months.Northern.Array, "North month array should match")
	assert.Equal(t, simple.Availability.MonthNorthern, actual.Months.Northern.Ranges, "North month ranges should match")
	assert.Equal(t, 2, len(actual.Months.Northern.Sequences), "North should have two sequences")
	// south
	assert.False(t, actual.Months.Southern.Always, "Should not be available all year in south")
	assert.Equal(t, simple.Availability.MonthArraySouthern, actual.Months.Southern.Array, "South month array should match")
	assert.Equal(t, simple.Availability.MonthSouthern, actual.Months.Southern.Ranges, "South month ranges should match")
	assert.Equal(t, 1, len(actual.Months.Southern.Sequences), "South should have one sequence")

	// times
	assert.True(t, actual.Times.Always, "Should not be available all day")
	assert.Equal(t, simple.Availability.Time, actual.Times.Text, "Time text array should match")
	assert.Equal(t, simple.Availability.TimeArray, actual.Times.Array, "Time array array should match")

}

func Test_has(t *testing.T) {
	simple := apistructures.Samples()["simple"]
	actual := meta(simple)
	// test has
	assert.True(t, actual.Has.Location)
	assert.True(t, actual.Has.Rarity)
	assert.True(t, actual.Has.Price)
	assert.True(t, actual.Has.Shadow)
	assert.False(t, actual.Has.Speed)
}
