package smfp_test

import (
	"testing"

	"github.com/devjaewon/smfp/v1"
	"github.com/stretchr/testify/assert"
)

type (
	ReduceGiven struct {
		Value int
	}
)

func TestReduce(t *testing.T) {
	t.Run("get result accumulating array element process", func(t *testing.T) {
		assert := assert.New(t)

		givens := []ReduceGiven{{1}, {2}, {3}, {4}, {5}}
		expected := 15

		result := smfp.Reduce(func(ac int, given ReduceGiven, index int) int {
			return ac + given.Value
		}, func() int {
			return 0
		}).Execute(givens).(int)

		assert.Equal(result, expected)
	})
}
