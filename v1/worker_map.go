package smfp

func Map[I, O any](mapper func(item I, index int) O) *Worker {
	return &Worker{
		task: func(input any) any {
			output := []O{}

			for i, item := range input.([]I) {
				output = append(output, mapper(item, i))
			}

			return output
		},
	}
}
