package smfp_test

import (
	"testing"

	"github.com/devjaewon/smfp/v1"
	"github.com/stretchr/testify/assert"
)

type (
	PipeGiven struct {
		Value int
	}
)

func TestPipe(t *testing.T) {
	t.Run("works multiple pipe", func(t *testing.T) {
		assert := assert.New(t)

		givens := []PipeGiven{{1}, {2}, {3}, {4}, {5}, {6}, {7}}
		expected := []int{15, 16, 17}

		results := smfp.Pipe[[]PipeGiven]([]*smfp.Worker{
			smfp.Map(func(item PipeGiven, index int) int {
				return item.Value + 10
			}),
			smfp.Filter(func(item int, index int) bool {
				return item > 14
			}),
		}).Execute(givens).([]int)

		assert.EqualValues(results, expected)
	})
}
