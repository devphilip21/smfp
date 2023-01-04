package smfp_test

import (
	"testing"

	"github.com/devjoeween/smfp/v1"
)

type (
	MapGiven struct {
		Value int
	}
)

func TestMap(t *testing.T) {
	t.Run("map array items A to B", func(t *testing.T) {
		givens := []MapGiven{{1}, {2}, {3}, {4}, {5}, {6}}
		expected := []int{11, 12, 13, 14, 15, 16}

		results := smfp.Map(func(given MapGiven, index int) int {
			return given.Value + 10
		}).Execute(givens)

		for i, result := range results.([]int) {
			if expected[i] != result {
				t.Fail()
				return
			}
		}
	})
}
