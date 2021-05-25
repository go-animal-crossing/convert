package targetstructures

import "time"

func (o *Output) addForListing(
	list *Listing,
	northNths []int,
	southNths []int,
	item Item,
	id string,
) {

	SetDefaultTypeMapsForListing(list)
	SetDefaultMonthMapsForListing(list)

	for _, mth := range northNths {
		month := time.Month(mth).String()
		list.Months[month][id] = item
	}
	for _, mth := range southNths {
		month := time.Month(mth).String()
		list.Months[month][id] = item
	}

	if item.Attributes.Type.Slug == "bugs" {
		list.Bugs[id] = item
	} else if item.Attributes.Type.Slug == "fish" {
		list.Fish[id] = item
	} else if item.Attributes.Type.Slug == "sea-creatures" {
		list.Sea[id] = item
	}

}
