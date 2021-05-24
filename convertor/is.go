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

func GenerateIs(t time.Time, item targetstructures.Item) targetstructures.Is {
	n := item.Attributes.Availability.Months.Northern
	north := targetstructures.IsHemisphere{
		New:       IsNew(t, n.Sequences, n.Always),
		Leaving:   IsLeaving(t, n.Sequences, n.Always),
		Available: IsAvailable(t, n.Array, n.Always),
	}
	s := item.Attributes.Availability.Months.Southern
	south := targetstructures.IsHemisphere{
		New:       IsNew(t, s.Sequences, s.Always),
		Leaving:   IsLeaving(t, s.Sequences, s.Always),
		Available: IsAvailable(t, s.Array, s.Always),
	}
	return targetstructures.Is{Northern: north, Southern: south}
}
