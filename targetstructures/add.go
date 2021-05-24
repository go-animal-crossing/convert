package targetstructures

import (
	"convert/util"
	"time"
)

func (o *Output) leaving(item Item, id string) {

	if item.Meta.Is.Northern.Leaving || item.Meta.Is.Southern.Leaving {
		o.Leaving.Current[id] = item

		if item.Meta.Is.Northern.Leaving {
			o.Leaving.Northern.Current[id] = item
			// add to leaving month
			sequence := item.Attributes.Availability.Months.Northern.Sequences
			lasts := util.NthOfSequences(sequence, -1)
			for _, mth := range lasts {
				month := time.Month(mth).String()
				if o.Leaving.Northern.Months[month] == nil {
					o.Leaving.Northern.Months[month] = make(map[string]Item)
				}
				o.Leaving.Northern.Months[month][id] = item
			}
		}
		if item.Meta.Is.Southern.Leaving {
			o.Leaving.Southern.Current[id] = item
			// add to leaving month
			sequence := item.Attributes.Availability.Months.Southern.Sequences
			lasts := util.NthOfSequences(sequence, -1)
			for _, mth := range lasts {
				month := time.Month(mth).String()
				if o.Leaving.Southern.Months[month] == nil {
					o.Leaving.Southern.Months[month] = make(map[string]Item)
				}
				o.Leaving.Southern.Months[month][id] = item
			}
		}

	}
}

func (o *Output) new(item Item, id string) {

	if item.Meta.Is.Northern.New || item.Meta.Is.Southern.New {
		o.New.Current[id] = item

		if item.Meta.Is.Northern.New {
			o.New.Northern.Current[id] = item
			// add to new month
			sequence := item.Attributes.Availability.Months.Northern.Sequences
			lasts := util.NthOfSequences(sequence, -1)
			for _, mth := range lasts {
				month := time.Month(mth).String()
				if o.New.Northern.Months[month] == nil {
					o.New.Northern.Months[month] = make(map[string]Item)
				}
				o.New.Northern.Months[month][id] = item
			}
		}
		if item.Meta.Is.Southern.New {
			o.New.Southern.Current[id] = item
			// add to new month
			sequence := item.Attributes.Availability.Months.Southern.Sequences
			lasts := util.NthOfSequences(sequence, -1)
			for _, mth := range lasts {
				month := time.Month(mth).String()
				if o.New.Southern.Months[month] == nil {
					o.New.Southern.Months[month] = make(map[string]Item)
				}
				o.New.Southern.Months[month][id] = item
			}
		}

	}
}

func (o *Output) available(item Item, id string) {

	if item.Meta.Is.Northern.Available || item.Meta.Is.Southern.Available {
		o.Available.Current[id] = item

		o.adder(
			item.Meta.Is.Northern.Available,
			&o.Available.Northern,
			item.Attributes.Availability.Months.Northern.Sequences,
			item,
			id,
		)

		o.adder(
			item.Meta.Is.Southern.Available,
			&o.Available.Southern,
			item.Attributes.Availability.Months.Southern.Sequences,
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
