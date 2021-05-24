package targetstructures

import "time"

func New() Output {
	return Output{
		Time: time.Now().UTC(),
		All:  make(map[string]Item),
		Bugs: make(map[string]Item),
		Fish: make(map[string]Item),
		Sea:  make(map[string]Item),
		Leaving: ItemTypeHemisphere{
			All: make(map[string]Item),
			Northern: TypedItems{
				All:  make(map[string]Item),
				Bugs: make(map[string]Item),
				Fish: make(map[string]Item),
				Sea:  make(map[string]Item),
			},
			Southern: TypedItems{
				All:  make(map[string]Item),
				Bugs: make(map[string]Item),
				Fish: make(map[string]Item),
				Sea:  make(map[string]Item),
			},
		},
		New: ItemTypeHemisphere{
			All: make(map[string]Item),
			Northern: TypedItems{
				All:  make(map[string]Item),
				Bugs: make(map[string]Item),
				Fish: make(map[string]Item),
				Sea:  make(map[string]Item),
			},
			Southern: TypedItems{
				All:  make(map[string]Item),
				Bugs: make(map[string]Item),
				Fish: make(map[string]Item),
				Sea:  make(map[string]Item),
			},
		},
		Available: ItemTypeHemisphere{
			All: make(map[string]Item),
			Northern: TypedItems{
				All:  make(map[string]Item),
				Bugs: make(map[string]Item),
				Fish: make(map[string]Item),
				Sea:  make(map[string]Item),
			},
			Southern: TypedItems{
				All:  make(map[string]Item),
				Bugs: make(map[string]Item),
				Fish: make(map[string]Item),
				Sea:  make(map[string]Item),
			},
		},
	}
}
