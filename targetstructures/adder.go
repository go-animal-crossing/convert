package targetstructures

import (
	"convert/util"
	"time"
)

func (o *Output) adder(
	is bool,
	hemisphere *ListingByHemisphere,
	sequence [][]int,
	item Item,
	id string,
) {

	if is {
		hemisphere.Current[id] = item
		// add to available month
		lasts := util.NthOfSequences(sequence, -1)
		for _, mth := range lasts {
			month := time.Month(mth).String()
			if hemisphere.Months[month] == nil {
				hemisphere.Months[month] = make(map[string]Item)
			}
			hemisphere.Months[month][id] = item
		}
	}

}
