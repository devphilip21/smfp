package smfp_test

import (
	"testing"

	"github.com/devjoeween/smfp/v1"
	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	t.Run("if even one element satisfies the condition, return true", func(t *testing.T) {
		assert := assert.New(t)
		given := []int{3, 4, 5, 6, 7, 8}

		result := smfp.Some(func(item int, index int) bool {
			return item == 6
		}).Execute(given).(bool)

		assert.Equal(result, true)
	})

	t.Run("if no one satisfies the condition, return false", func(t *testing.T) {
		assert := assert.New(t)
		given := []int{2, 5, 6, 7, 8}

		result := smfp.Some(func(item int, index int) bool {
			return item < 2 || item > 8
		}).Execute(given).(bool)

		assert.Equal(result, false)
	})
}
