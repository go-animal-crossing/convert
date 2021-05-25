package targetstructures

func (o *Output) Add(item Item) {
	id := item.ID
	o.All[id] = item
}
