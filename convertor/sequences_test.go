package convertor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateSequences(t *testing.T) {

	// [1,2,3] => 0:[1,2,3]
	nums := []int{1, 2, 3}
	res := GenerateSequences(nums)
	// should have 1 sequence
	assert.Equal(t, 1, len(res))
	// that sequence should have 3 items
	assert.Equal(t, 3, len(res[0]))

	// [11, 12, 1, 2] => 0:[11,12,1,2]
	nums = []int{11, 12, 1, 2}
	res = GenerateSequences(nums)
	// should have 1 sequence
	assert.Equal(t, 1, len(res))
	// that sequence should have 4 items
	assert.Equal(t, 4, len(res[0]))

	// [1,2,5,6] => 0:[1,2], 1:[5,6]
	nums = []int{1, 2, 5, 6}
	res = GenerateSequences(nums)
	// should have 1 sequence
	assert.Equal(t, 2, len(res))
	// both should have length of 2
	assert.Equal(t, 2, len(res[0]))
	assert.Equal(t, 2, len(res[1]))

	// [11,12,1,5,6] => 0:[11,12,1], 1:[5,6]
	nums = []int{11, 12, 1, 5, 6}
	res = GenerateSequences(nums)
	assert.Equal(t, 2, len(res))
	// sequences should have length of 3 and 2
	assert.Equal(t, 3, len(res[0]))
	assert.Equal(t, 2, len(res[1]))

	// [1,2, 11,12,1, 5] => 0[1,2], 1:[11,12,1], 2:[5]
	nums = []int{1, 2, 11, 12, 1, 5}
	res = GenerateSequences(nums)
	assert.Equal(t, 3, len(res))
	// sequences should have length of 3 and 2
	assert.Equal(t, 2, len(res[0]))
	assert.Equal(t, 3, len(res[1]))
	assert.Equal(t, 1, len(res[2]))

	// [12,1,2,3, 1, 6,7, 10,11,12] => 0:[12,1,2,3], 1:[1], 2:[6,7], 3:[10,11,12]
	nums = []int{12, 1, 2, 3, 1, 6, 7, 10, 11, 12}
	res = GenerateSequences(nums)
	assert.Equal(t, 4, len(res))
	// sequences should have length of 3 and 2
	assert.Equal(t, 4, len(res[0]))
	assert.Equal(t, 1, len(res[1]))
	assert.Equal(t, 2, len(res[2]))
	assert.Equal(t, 3, len(res[3]))

}

func Test_SequenceToText(t *testing.T) {

	// Feb
	nums := []int{2}
	seq := GenerateSequences(nums)
	res := GenerateSequenceText(seq)
	assert.Equal(t, "February", res)

	nums = []int{1, 3}
	seq = GenerateSequences(nums)
	res = GenerateSequenceText(seq)
	assert.Equal(t, "January, March", res)

	// [1,2,3] => 0:[1,2,3] => Jan -> Mar
	nums = []int{1, 2, 3}
	seq = GenerateSequences(nums)
	res = GenerateSequenceText(seq)
	assert.Equal(t, "January - March", res)

	nums = []int{1, 2, 4, 5}
	seq = GenerateSequences(nums)
	res = GenerateSequenceText(seq)
	assert.Equal(t, "January - February, April - May", res)

	nums = []int{11, 12, 1, 4, 5}
	seq = GenerateSequences(nums)
	res = GenerateSequenceText(seq)
	assert.Equal(t, "November - January, April - May", res)

	nums = []int{11, 12, 1, 4, 5, 11, 12}
	seq = GenerateSequences(nums)
	res = GenerateSequenceText(seq)
	assert.Equal(t, "November - January, April - May, November - December", res)

	nums = []int{11, 12, 1, 4, 5, 12}
	seq = GenerateSequences(nums)
	res = GenerateSequenceText(seq)
	assert.Equal(t, "November - January, April - May, December", res)

}
