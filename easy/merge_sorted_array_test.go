package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedArray(t *testing.T) {
	for i, tt := range []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		res   []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			res:   []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			res:   []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			res:   []int{1},
		},
	} {
		i, tt := i, tt
		t.Run(fmt.Sprintf("Run %d", i), func(t *testing.T) {
			MergeSortedArray(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.res, tt.nums1)
		})
	}
}
