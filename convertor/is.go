package convertor

import (
	"convert/targetstructures"
	"convert/util"
	"time"
)

func IsNew(t time.Time, sequences [][]int, always bool) bool {
	if always {
		return false
	}
	month := int(t.Month())
	ns := util.NthOfSequences(sequences, 0)
	return util.Contains(ns, month)
}

func IsLeaving(t time.Time, sequences [][]int, always bool) bool {
	if always {
		return false
	}
	month := int(t.Month())
	ns := util.NthOfSequences(sequences, -1)
	return util.Contains(ns, month)
}

func IsAvailable(t time.Time, sequences []int, always bool) bool {
	if always {
		return true
	}
	month := int(t.Month())
	return util.Contains(sequences, month)
}

func GenerateIs(t time.Time, item targetstructures.Item) map[string]map[string]bool {
	n := item.Attributes.Availability.Months.Northern
	s := item.Attributes.Availability.Months.Southern

	is := make(map[string]map[string]bool)

	is["northern"] = make(map[string]bool)
	is["northern"]["new"] = IsNew(t, n.Sequences, n.Always)
	is["northern"]["leaving"] = IsLeaving(t, n.Sequences, n.Always)
	is["northern"]["available"] = IsAvailable(t, n.Array, n.Always)

	is["southern"] = make(map[string]bool)
	is["southern"]["new"] = IsNew(t, s.Sequences, s.Always)
	is["southern"]["leaving"] = IsLeaving(t, s.Sequences, s.Always)
	is["southern"]["available"] = IsAvailable(t, s.Array, s.Always)

	return is
}
