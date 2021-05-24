package targetstructures

import "time"

func New() Output {
	return Output{
		Time: time.Now().UTC(),
		All:  make(map[string]Item),
		Bugs: make(map[string]Item),
		Fish: make(map[string]Item),
		Sea:  make(map[string]Item),
		Leaving: Listing{
			Current: make(map[string]Item),
			Months:  make(map[string]map[string]Item),
			Northern: ListingByHemisphere{
				Current: make(map[string]Item),
				Months:  make(map[string]map[string]Item),
			},
			Southern: ListingByHemisphere{
				Current: make(map[string]Item),
				Months:  make(map[string]map[string]Item),
			},
		},
		New: Listing{
			Current: make(map[string]Item),
			Months:  make(map[string]map[string]Item),
			Northern: ListingByHemisphere{
				Current: make(map[string]Item),
				Months:  make(map[string]map[string]Item),
			},
			Southern: ListingByHemisphere{
				Current: make(map[string]Item),
				Months:  make(map[string]map[string]Item),
			},
		},
		Available: Listing{
			Current: make(map[string]Item),
			Months:  make(map[string]map[string]Item),
			Northern: ListingByHemisphere{
				Current: make(map[string]Item),
				Months:  make(map[string]map[string]Item),
			},
			Southern: ListingByHemisphere{
				Current: make(map[string]Item),
				Months:  make(map[string]map[string]Item),
			},
		},
	}
}
