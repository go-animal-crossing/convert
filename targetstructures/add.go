package targetstructures

import "convert/util"

func (o *Output) addForListAndHemispheres(
	list *Listing,
	north *ListingByHemisphere,
	south *ListingByHemisphere,
	nIs bool,
	sIs bool,
	nNths []int,
	sNths []int,
	item Item,
	id string,
) {
	o.addForListing(list, nNths, sNths, item, id)
	o.addForHemisphere(nIs, north, nNths, item, id)
	o.addForHemisphere(sIs, south, sNths, item, id)

}

func (o *Output) leaving(item Item, id string) {

	if item.Meta.Is.Northern.Leaving || item.Meta.Is.Southern.Leaving {
		o.Leaving.Current[id] = item
		n := item.Attributes.Availability.Months.Northern.Sequences
		nNths := util.NthOfSequences(n, -1)

		s := item.Attributes.Availability.Months.Southern.Sequences
		sNths := util.NthOfSequences(s, -1)

		o.addForListAndHemispheres(
			&o.Leaving,
			&o.Leaving.Northern,
			&o.Leaving.Southern,
			item.Meta.Is.Northern.Leaving,
			item.Meta.Is.Southern.Leaving,
			nNths,
			sNths,
			item,
			id,
		)

	}
}

func (o *Output) new(item Item, id string) {

	if item.Meta.Is.Northern.New || item.Meta.Is.Southern.New {
		o.New.Current[id] = item
		n := item.Attributes.Availability.Months.Northern.Sequences
		nNths := util.NthOfSequences(n, -0)
		s := item.Attributes.Availability.Months.Southern.Sequences
		sNths := util.NthOfSequences(s, -0)

		o.addForListAndHemispheres(
			&o.New,
			&o.New.Northern,
			&o.New.Southern,
			item.Meta.Is.Northern.New,
			item.Meta.Is.Southern.New,
			nNths,
			sNths,
			item,
			id,
		)

	}
}

func (o *Output) available(item Item, id string) {

	if item.Meta.Is.Northern.Available || item.Meta.Is.Southern.Available {
		o.Available.Current[id] = item
		n := item.Attributes.Availability.Months.Northern.Array
		s := item.Attributes.Availability.Months.Southern.Array

		o.addForListAndHemispheres(
			&o.Available,
			&o.Available.Northern,
			&o.Available.Southern,
			item.Meta.Is.Northern.Available,
			item.Meta.Is.Southern.Available,
			n,
			s,
			item,
			id,
		)

	}
}

func (o *Output) Add(item Item) {
	id := item.ID
	o.All[id] = item

	if item.Attributes.Type.Slug == "bugs" {
		o.Bugs[id] = item
	} else if item.Attributes.Type.Slug == "fish" {
		o.Fish[id] = item
	} else if item.Attributes.Type.Slug == "sea-creatures" {
		o.Sea[id] = item
	}

	o.leaving(item, id)
	o.new(item, id)
	o.available(item, id)
}
