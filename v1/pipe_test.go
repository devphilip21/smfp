package smfp_test

import (
	"testing"

	"github.com/devjoeween/smfp/v1"
)

type (
	PipeGiven struct {
		Value int
	}
)

func TestPipe(t *testing.T) {
	t.Run("works multiple pipe", func(t *testing.T) {
		givens := []PipeGiven{
			{1}, {2}, {3}, {4}, {5}, {6}, {7},
		}
		expected := []int{11, 12, 13, 14, 15, 16, 17}

		results := smfp.Pipe[[]PipeGiven]([]*smfp.Worker{
			smfp.Map(func(item PipeGiven, index int) int {
				return item.Value + 10
			}),
			smfp.Map(func(item int, index int) PipeGiven {
				return PipeGiven{
					Value: item,
				}
			}),
		}).Execute(givens).([]PipeGiven)

		for i, result := range results {
			if expected[i] != result.Value {
				t.Fail()
				return
			}
		}
	})
}
