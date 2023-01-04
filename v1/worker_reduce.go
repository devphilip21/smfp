package smfp

func Reduce[I, O any](
	reducer func(accumulator O, item I, index int) O,
	initializer func() O,
) *Worker {
	return &Worker{
		task: func(input any) any {
			var result O = initializer()
			items := input.([]I)

			for i, item := range items {
				result = reducer(result, item, i)
			}

			return result
		},
	}
}
