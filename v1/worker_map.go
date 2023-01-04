package smfp

func Map[I, O any](mapper func(item I, index int) O) *Worker {
	return &Worker{
		task: func(input any) any {
			output := []O{}

			for i, input := range input.([]I) {
				output = append(output, mapper(input, i))
			}

			return output
		},
	}
}
