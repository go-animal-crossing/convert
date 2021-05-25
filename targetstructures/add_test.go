package targetstructures

import (
	"convert/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_simple(t *testing.T) {
	r := New()
	// simple is a bug and available all the time
	simple := Samples()["simple"][0]
	r.Add(simple)

	assert.Equal(t, 1, len(r.All))

	// get all tags
	tags := []string{}
	for _, i := range r.All {
		tags = append(tags, i.Tags...)
	}

	assert.Contains(t, tags, "type_bugs")
	assert.NotContains(t, tags, "type_fish")
	assert.NotContains(t, tags, "type_sea_creatures")
	assert.NotContains(t, tags, "leaving_northern_may")

}

func Test_Add_multi(t *testing.T) {
	r := New()
	multi := Samples()["multi"]

	for _, i := range multi {
		r.Add(i)
	}
	assert.Equal(t, 5, len(r.All))

	// get all tags
	tags := []string{}
	for _, i := range r.All {
		tags = append(tags, i.Tags...)
	}
	//fmt.Printf("%v\n", tags)

	assert.Contains(t, tags, "type_bugs")
	count := util.Count(tags, "type_bugs")
	assert.Equal(t, 1, count)

	assert.Contains(t, tags, "type_fish")
	count = util.Count(tags, "type_fish")
	assert.Equal(t, 4, count)

	assert.NotContains(t, tags, "type_sea_creatures")
	assert.NotContains(t, tags, "leaving_may_northern")

	assert.Contains(t, tags, "leaving_may_southern")
	count = util.Count(tags, "leaving_may_southern")
	assert.Equal(t, 1, count)

	assert.Contains(t, tags, "new_may")
	assert.Contains(t, tags, "new_may_northern")
	count = util.Count(tags, "new_may_northern")
	assert.Equal(t, 1, count)

	// 2 new fish in may - bitterfish in south, rainbow in north
	assert.Contains(t, tags, "type_fish_new_may")
	count = util.Count(tags, "type_fish_new_may")
	assert.Equal(t, 2, count)

	// just one new fish in south for may - bitterling
	assert.Contains(t, tags, "type_fish_new_may_southern")
	count = util.Count(tags, "type_fish_new_may_southern")
	assert.Equal(t, 1, count)

	// rainbow fish
	assert.Contains(t, tags, "type_fish_new_may_northern")
	count = util.Count(tags, "type_fish_new_may_northern")
	assert.Equal(t, 1, count)

	assert.Contains(t, tags, "leaving_december_southern")
	count = util.Count(tags, "leaving_december_southern")
	assert.Equal(t, 1, count)

	assert.Contains(t, tags, "leaving_december")
	count = util.Count(tags, "leaving_december")
	assert.Equal(t, 1, count)

	assert.Contains(t, tags, "type_fish_leaving_december_southern")
	count = util.Count(tags, "type_fish_leaving_december_southern")
	assert.Equal(t, 1, count)
}
