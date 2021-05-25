package targetstructures

import (
	"time"
)

func (o *Output) addForHemisphere(
	is bool,
	hemisphere *ListingByHemisphere,
	nths []int,
	item Item,
	id string,
) {

	SetDefaultTypeMapsForHemisphere(hemisphere)
	SetDefaultMonthMapsForHemisphere(hemisphere)

	if is {
		hemisphere.Current[id] = item
		for _, mth := range nths {
			month := time.Month(mth).String()
			hemisphere.Months[month][id] = item
		}

		if item.Attributes.Type.Slug == "bugs" {
			hemisphere.Bugs[id] = item
		} else if item.Attributes.Type.Slug == "fish" {
			hemisphere.Fish[id] = item
		} else if item.Attributes.Type.Slug == "sea-creatures" {
			hemisphere.Sea[id] = item
		}

	}

}
