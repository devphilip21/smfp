package smfp_test

import (
	"testing"

	"github.com/devjoeween/smfp/v1"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("filter elements in array", func(t *testing.T) {
		assert := assert.New(t)

		givens := []int{1, 2, 3, 4, 5, 6, 7}
		expected := []int{3, 4, 5}

		results := smfp.Filter(func(num int, index int) bool {
			return num > 2 && num < 6
		}).Execute(givens).([]int)

		assert.EqualValues(results, expected)
	})
}
