package targetstructures

import (
	"testing"
)

func Test_Add_simple(t *testing.T) {
	r := New()
	// simple is a bug and available all the time
	simple := Samples()["simple"][0]
	//simple.Tags = convertor.Tags(simple)

	r.Add(simple)

	// fmt.Printf("%v", r.All["00cb41a2dcac815286129ab947d02a3f2a4b79ce"].Tags)
	// assert.Equal(t, 1, len(r.All))
	// assert.Equal(t, 1, len(r.Bugs))
	// assert.Equal(t, 0, len(r.Fish))
	// assert.Equal(t, 0, len(r.Sea))

	// assert.Equal(t, 0, len(r.Leaving.Northern.Current))

	// assert.Equal(t, 0, len(r.Leaving.Southern.Current))
	// assert.Equal(t, 0, len(r.New.Northern.Current))
	// assert.Equal(t, 0, len(r.New.Southern.Current))
	// assert.Equal(t, 0, len(r.Available.Northern.Current))
	// assert.Equal(t, 1, len(r.Available.Southern.Current))

	// assert.Equal(t, 1, len(r.Available.Bugs))
	// assert.Equal(t, 1, len(r.Available.Months["May"]))

}

func Test_Add_multi(t *testing.T) {
	r := New()
	multi := Samples()["multi"]

	for _, i := range multi {
		r.Add(i)
	}

	// assert.Equal(t, 5, len(r.All))
	// assert.Equal(t, 1, len(r.Bugs))
	// assert.Equal(t, 4, len(r.Fish))

	// assert.Equal(t, 0, len(r.Leaving.Northern.Current))
	// assert.Equal(t, 1, len(r.Leaving.Southern.Current))

	// assert.Equal(t, 1, len(r.New.Northern.Current))
	// assert.Equal(t, 0, len(r.New.Northern.Bugs))

	// assert.Equal(t, 1, len(r.New.Southern.Current))
	// assert.Equal(t, 1, len(r.New.Southern.Fish))
	// //
	// assert.Equal(t, 1, len(r.Leaving.Southern.Months["December"]))
	// assert.Equal(t, 1, len(r.Leaving.Months["December"]))

	// assert.Equal(t, 1, len(r.Leaving.Southern.Fish))
}
