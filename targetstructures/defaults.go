package targetstructures

import (
	"time"
)

func SetDefaultTypeMapsForHemisphere(hemisphere *ListingByHemisphere) {
	if hemisphere.Bugs == nil {
		hemisphere.Bugs = make(map[string]Item)
	}
	if hemisphere.Fish == nil {
		hemisphere.Fish = make(map[string]Item)
	}
	if hemisphere.Sea == nil {
		hemisphere.Sea = make(map[string]Item)
	}
}

func SetDefaultMonthMapsForHemisphere(hemisphere *ListingByHemisphere) {
	for i := 1; i <= 12; i++ {
		month := time.Month(i).String()
		if hemisphere.Months[month] == nil {
			hemisphere.Months[month] = make(map[string]Item)
		}
	}
}

//
func SetDefaultTypeMapsForListing(list *Listing) {
	if list.Bugs == nil {
		list.Bugs = make(map[string]Item)
	}
	if list.Fish == nil {
		list.Fish = make(map[string]Item)
	}
	if list.Sea == nil {
		list.Sea = make(map[string]Item)
	}
}
func SetDefaultMonthMapsForListing(list *Listing) {
	for i := 1; i <= 12; i++ {
		month := time.Month(i).String()
		if list.Months[month] == nil {
			list.Months[month] = make(map[string]Item)
		}
	}
}
