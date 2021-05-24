package targetstructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_simple(t *testing.T) {
	r := New()
	// simple is a bug and available all the time
	simple := Samples()["simple"][0]

	r.Add(simple)

	assert.Equal(t, len(r.All), 1)
	assert.Equal(t, len(r.Bugs), 1)
	assert.Equal(t, len(r.Fish), 0)
	assert.Equal(t, len(r.Sea), 0)

	assert.Equal(t, len(r.Leaving.Northern.Current), 0)

	assert.Equal(t, len(r.Leaving.Southern.Current), 0)
	assert.Equal(t, len(r.New.Northern.Current), 0)
	assert.Equal(t, len(r.New.Southern.Current), 0)
	assert.Equal(t, len(r.Available.Northern.Current), 0)
	assert.Equal(t, len(r.Available.Southern.Current), 1)

}

// func Test_Add_multi(t *testing.T) {
// 	r := New()
// 	multi := Samples()["multi"]

// 	for _, i := range multi {
// 		r.Add(i)
// 	}

// 	assert.Equal(t, len(r.All), 4)
// 	assert.Equal(t, len(r.Bugs), 1)
// 	assert.Equal(t, len(r.Fish), 3)

// 	// assert.Equal(t, len(r.Leaving.Northern.All), 0)
// 	// assert.Equal(t, len(r.Leaving.Southern.All), 0)

// 	// assert.Equal(t, len(r.New.Northern.All), 1)
// 	// assert.Equal(t, len(r.New.Northern.Fish), 1)
// 	// assert.Equal(t, len(r.New.Southern.All), 1)
// 	// assert.Equal(t, len(r.New.Southern.Fish), 1)
// }
