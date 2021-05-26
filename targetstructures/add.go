package targetstructures

func (o *Output) Add(item Item) {
	id := item.ID
	o.Sorted = append(o.Sorted, item)
	o.All[id] = item
}
