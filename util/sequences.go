package util

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
