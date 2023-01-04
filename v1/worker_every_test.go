package smfp_test

import (
	"testing"

	"github.com/devjoeween/smfp/v1"
	"github.com/stretchr/testify/assert"
)

func TestEvery(t *testing.T) {
	t.Run("every elements satisfy the condition, return true", func(t *testing.T) {
		assert := assert.New(t)
		given := []int{3, 4, 5, 6, 7, 8}

		result := smfp.Every(func(item int, index int) bool {
			return item > 2 && item < 9
		}).Execute(given).(bool)

		assert.Equal(result, true)
	})

	t.Run("if any of the conditions are not satisfied, return false", func(t *testing.T) {
		assert := assert.New(t)
		given := []int{2, 5, 6, 7, 8}

		result := smfp.Every(func(item int, index int) bool {
			return item > 2 && item < 9
		}).Execute(given).(bool)

		assert.Equal(result, false)
	})
}
