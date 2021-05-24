package convertor

import (
	"fmt"
)

func GenerateSequences(numbers []int) [][]int {

	max := 12
	current := numbers[0] - 1
	sequence := make([]int, 0)
	sequences := make([][]int, 0)
	// what doing here is looping over set of ints and appending
	// to a slice of ints
	for i := 0; i < len(numbers); i++ {
		num := numbers[i]
		test := current + 1
		//fmt.Printf("[%v - %v - %v]\n", num, current, test)
		// this triggers the next index in the main slice by appending and clearing current
		if test != num {
			sequences = append(sequences, sequence)
			//fmt.Printf("sequence complete: \n%v\n%v\n", sequence, sequences)
			//fmt.Printf("next sequence starting: %v\n", sequence)
			sequence = make([]int, 0)
		}

		sequence = append(sequence, num)
		current = num % max
	}
	// append the last sequence
	sequences = append(sequences, sequence)

	return sequences
}

func GenerateSequenceText(sequences [][]int) string {

	dateString := ""
	// date format to use for outputting string (currently full month name)
	f := DateFormat()

	for i, sequence := range sequences {
		// first item in the slice
		// - grab and generate a Time
		first := sequence[0]
		firstDate := MonthToTime(first)
		// always have a "," at the start, unless this is the first item
		prefix := ", "
		if i == 0 {
			prefix = ""
		}

		dateString = fmt.Sprintf("%s%s%s", dateString, prefix, firstDate.Format(f))

		l := len(sequence)

		if l > 1 {
			last := sequence[l-1]
			lastDate := MonthToTime(last)
			dateString = fmt.Sprintf("%s - %s", dateString, lastDate.Format(f))
		}

	}
	//fmt.Printf("%v => %v\n", sequences, dateString)
	return dateString

}

// return the nth index of each sequence
func NthOfSequences(sequences [][]int, n int) (ns []int) {
	for _, seq := range sequences {
		check := n
		if n == -1 {
			check = len(seq) - 1
		}
		ns = append(ns, seq[check])
	}
	return
}
