package targetstructures

import "sort"

func (o *Output) Sort() {

	sort.Slice(o.Sorted, func(i, j int) bool {
		return o.Sorted[i].Attributes.Titles.Safe < o.Sorted[j].Attributes.Titles.Safe
	})
}
