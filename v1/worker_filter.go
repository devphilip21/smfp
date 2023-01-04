package smfp

func Filter[I any](predecate func(item I, index int) (leave bool)) *Worker {
	return &Worker{
		task: func(input any) any {
			output := []I{}

			for i, item := range input.([]I) {
				if predecate(item, i) {
					output = append(output, item)
				}
			}

			return output
		},
	}
}
