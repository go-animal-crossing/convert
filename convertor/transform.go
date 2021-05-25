package convertor

import (
	"convert/apistructures"
	"convert/targetstructures"
	"convert/util"
	"fmt"
	"strings"
	"time"
)

func titles(item apistructures.Item) targetstructures.Safe {
	return targetstructures.Safe{
		Original: item.Names.EuEn,
		Safe:     util.Title(util.Safe(item.Names.EuEn)),
	}
}

func uris(item apistructures.Item) targetstructures.Uris {
	t := typeMeta[item.Type]
	return targetstructures.Uris{
		URL:  util.URL(t.Slug, util.Safe(item.Names.EuEn)),
		Slug: util.Slugify(util.Safe(item.Names.EuEn)),
	}
}

func prices(item apistructures.Item) targetstructures.Prices {
	return targetstructures.Prices{
		Store: item.Price,
		Cj:    item.PriceCj,
		Flick: item.PriceFlick,
	}
}

func phrases(item apistructures.Item) targetstructures.Phrases {
	capture := targetstructures.Safe{
		Original: item.CatchPhrase,
		Safe:     util.Safe(item.CatchPhrase),
	}
	museum := targetstructures.Safe{
		Original: item.MuseumPhrase,
		Safe:     util.Safe(item.MuseumPhrase),
	}
	return targetstructures.Phrases{
		Capture: capture,
		Museum:  museum,
	}
}

func images(item apistructures.Item) targetstructures.Images {
	thumb := targetstructures.Image{
		Direct: item.IconURI,
		Local:  util.ImagePath(item.Type, "thumb", item.FileName, "png"),
	}
	main := targetstructures.Image{
		Direct: item.ImageURI,
		Local:  util.ImagePath(item.Type, "main", item.FileName, "png"),
	}
	return targetstructures.Images{
		Thumb: thumb,
		Main:  main,
	}
}

func availability(item apistructures.Item) targetstructures.Availability {

	north := targetstructures.Hemisphere{
		Always: item.Availability.IsAllYear,
		Ranges: item.Availability.MonthNorthern,
		Array:  item.Availability.MonthArrayNorthern,
	}
	north.Sequences = GenerateSequences(north.Array)
	north.Text = GenerateSequenceText(north.Sequences)

	south := targetstructures.Hemisphere{
		Always: item.Availability.IsAllYear,
		Ranges: item.Availability.MonthSouthern,
		Array:  item.Availability.MonthArraySouthern,
	}
	south.Sequences = GenerateSequences(south.Array)
	south.Text = GenerateSequenceText(south.Sequences)

	months := targetstructures.Months{
		Always:   item.Availability.IsAllYear,
		Northern: north,
		Southern: south,
	}
	times := targetstructures.Times{
		Always: item.Availability.IsAllDay,
		Text:   item.Availability.Time,
		Array:  item.Availability.TimeArray,
	}
	return targetstructures.Availability{
		Location: item.Availability.Location,
		Rarity:   item.Availability.Rarity,
		Months:   months,
		Times:    times,
	}
}

func has(item apistructures.Item) targetstructures.Has {

	available := ((len(item.Availability.MonthArrayNorthern) > 0) ||
		(len(item.Availability.MonthArraySouthern) > 0))

	return targetstructures.Has{
		Price:        (item.Price > 0),
		Shadow:       (len(item.Shadow) > 0),
		Speed:        (len(item.Speed) > 0),
		Rarity:       (len(item.Availability.Rarity) > 0),
		Location:     (len(item.Availability.Location) > 0),
		Availability: available,
	}

}

func attributes(item apistructures.Item) targetstructures.Attributes {
	return targetstructures.Attributes{
		Type:         typeMeta[item.Type],
		Shadow:       item.Shadow,
		Speed:        item.Speed,
		Titles:       titles(item),
		URIS:         uris(item),
		Prices:       prices(item),
		Phrases:      phrases(item),
		Images:       images(item),
		Availability: availability(item),
	}
}

func Tags(item targetstructures.Item) []string {
	tags := []string{fmt.Sprintf("type_%s", item.Attributes.Type.Slug)}

	if item.Attributes.Availability.Months.Always {
		tags = append(tags, "available_always")
	}
	if item.Attributes.Availability.Months.Northern.Always {
		tags = append(tags, "available_northern_always")
	}
	if item.Attributes.Availability.Months.Southern.Always {
		tags = append(tags, "available_southern_always")
	}
	// get the Is data per month and generate tags for each
	for m := 1; m <= 12; m++ {
		month := time.Date(2021, time.Month(m), 1, 1, 0, 0, 0, time.UTC)
		is := GenerateIs(month, item)
		for hemiName, hemiData := range is {
			for prop, val := range hemiData {
				m := month.Month().String()
				tag := fmt.Sprintf("%s_%s", prop, m)
				htag := fmt.Sprintf("%s_%s_%s", prop, hemiName, m)
				if val {
					tags = append(tags, strings.ToLower(htag), strings.ToLower(tag))
				}
			}
		}
	}

	return tags
}

func Transform(item apistructures.Item) targetstructures.Item {
	target := targetstructures.Item{
		ID:         item.ID(),
		Attributes: attributes(item),
		Has:        has(item),
		Converted:  true,
	}
	target.Tags = Tags(target)

	return target

}
