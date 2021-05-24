package targetstructures

func typedAdd(data *TypedItems, item Item, id string, typeName string) {
	data.All[id] = item
	switch t := typeName; t {
	case "bugs":
		data.Bugs[id] = item
	case "fish":
		data.Fish[id] = item
	case "sea-creatures":
		data.Sea[id] = item
	}
}

func (o *Output) addToType(item Item, id string, itemType string) {
	// add based on the item type
	switch t := itemType; t {
	case "bugs":
		o.Bugs[id] = item
	case "fish":
		o.Fish[id] = item
	case "sea-creatures":
		o.Sea[id] = item
	}
}

func (o *Output) addToLeaving(item Item, id string, itemType string) {
	if item.Meta.Is.Northern.Leaving || item.Meta.Is.Southern.Leaving {
		o.Leaving.All[id] = item
	}
	if item.Meta.Is.Northern.Leaving {
		typedAdd(&o.Leaving.Northern, item, id, itemType)
	}
	if item.Meta.Is.Southern.Leaving {
		typedAdd(&o.Leaving.Southern, item, id, itemType)
	}
}

func (o *Output) addToNew(item Item, id string, itemType string) {
	if item.Meta.Is.Northern.New || item.Meta.Is.Southern.New {
		o.New.All[id] = item
	}

	if item.Meta.Is.Northern.New {
		typedAdd(&o.New.Northern, item, id, itemType)
	}
	if item.Meta.Is.Southern.New {
		typedAdd(&o.New.Southern, item, id, itemType)
	}
}

func (o *Output) addToAvailable(item Item, id string, itemType string) {
	if item.Meta.Is.Northern.Available || item.Meta.Is.Southern.Available {
		o.Available.All[id] = item
	}

	if item.Meta.Is.Northern.Available {
		typedAdd(&o.Available.Northern, item, id, itemType)
	}
	if item.Meta.Is.Southern.Available {
		typedAdd(&o.Available.Southern, item, id, itemType)
	}
}

func (o *Output) Add(item Item) {
	id := item.ID
	itemType := item.Attributes.Type.Slug

	o.All[id] = item
	o.addToType(item, id, itemType)
	o.addToLeaving(item, id, itemType)
	o.addToNew(item, id, itemType)
	o.addToAvailable(item, id, itemType)
}
