package utils_test

import (
	// std
	"testing"

	// 3rd-party libs
	"github.com/stretchr/testify/assert"

	// cess libs
	"github.com/CESSProject/cess-go-sdk/core/utils"
)

// Test cases for [TestNumsToByteStr] function, testing [utils.NumsToByteStr].
// It is an array of tuple (input, opts, expected result).
var numsToByteStrTests = []struct {
	input  []uint32
	opts   map[string]bool
	expect string
}{
	{[]uint32{15, 17}, map[string]bool{}, "0F11"},
	{[]uint32{15, 17}, map[string]bool{"prefix": true, "uppercase": false}, "0x0f11"},
	{[]uint32{15, 17}, map[string]bool{"space": true}, "0F 11"},
	{[]uint32{250, 15, 137, 153, 21, 63, 234, 235, 33, 164, 154, 189, 40, 122, 44, 96, 137, 183, 184, 180, 31, 212, 6, 2, 142, 22, 148, 22, 126, 164, 163, 105}, map[string]bool{}, "FA0F8999153FEAEB21A49ABD287A2C6089B7B8B41FD406028E1694167EA4A369"},
}

// Run through the test cases of [numsToByteStrTests] and check the result.
func TestNumsToByteStr(t *testing.T) {
	for _, tc := range numsToByteStrTests {
		res, err := utils.NumsToByteStr(tc.input, tc.opts)
		assert.Equal(t, tc.expect, res)
		assert.NoError(t, err)
	}

	// Test for error condition
	_, err := utils.NumsToByteStr([]uint32{256}, map[string]bool{})
	assert.Error(t, err)
}
