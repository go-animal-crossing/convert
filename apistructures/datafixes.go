package apistructures

func DataFixes(i Item) Item {
	// abalone has an issue with its southern data
	if i.FileName == "abalone" {
		i.Availability.MonthArraySouthern = []int{12, 1, 2, 3, 4, 5, 6, 7}
	}
	return i
}
