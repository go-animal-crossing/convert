package convertor

import (
	"convert/apistructures"
	"convert/targetstructures"
	"time"
)

func titles(item apistructures.Item) targetstructures.Safe {
	return targetstructures.Safe{
		Original: item.Names.EuEn,
		Safe:     Title(Safe(item.Names.EuEn)),
	}
}

func uris(item apistructures.Item) targetstructures.Uris {
	return targetstructures.Uris{
		URL:  URL(item.Type, Safe(item.Names.EuEn)),
		Slug: Slugify(Safe(item.Names.EuEn)),
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
		Safe:     Safe(item.CatchPhrase),
	}
	museum := targetstructures.Safe{
		Original: item.MuseumPhrase,
		Safe:     Safe(item.MuseumPhrase),
	}
	return targetstructures.Phrases{
		Capture: capture,
		Museum:  museum,
	}
}

func images(item apistructures.Item) targetstructures.Images {
	thumb := targetstructures.Image{
		Direct: item.IconURI,
		Local:  ImagePath(item.Type, "thumb", item.FileName, "png"),
	}
	main := targetstructures.Image{
		Direct: item.ImageURI,
		Local:  ImagePath(item.Type, "main", item.FileName, "png"),
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

	south := targetstructures.Hemisphere{
		Always: item.Availability.IsAllYear,
		Ranges: item.Availability.MonthSouthern,
		Array:  item.Availability.MonthArraySouthern,
	}
	south.Sequences = GenerateSequences(south.Array)

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

func meta(item apistructures.Item) targetstructures.Meta {
	available := ((len(item.Availability.MonthArrayNorthern) > 0) ||
		(len(item.Availability.MonthArraySouthern) > 0))

	has := targetstructures.Has{
		Price:        (item.Price > 0),
		Shadow:       (len(item.Shadow) > 0),
		Speed:        (len(item.Speed) > 0),
		Rarity:       (len(item.Availability.Rarity) > 0),
		Location:     (len(item.Availability.Location) > 0),
		Availability: available,
	}

	return targetstructures.Meta{
		Time: time.Now().UTC(),
		Has:  has,
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
func transform(item apistructures.Item) targetstructures.Item {

	target := targetstructures.Item{
		ID:         item.ID(),
		Attributes: attributes(item),
		Meta:       meta(item),
		Converted:  true,
	}

	return target

}
