package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type Answer []int
	tests := []struct {
		nums    []int
		target  int
		answers []Answer
	}{
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 6,
			answers: []Answer{
				{0, 4}, {1, 3},
			},
		},
	}
	for i, tt := range tests {
		assert.Contains(t, (tt.answers), TwoSum(tt.nums, tt.target), "", i, "test solution")
	}
}
