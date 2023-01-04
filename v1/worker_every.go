package smfp

func Every[I any](predicate func(item I, index int) bool) *Worker {
	return &Worker{
		task: func(input any) any {
			for i, item := range input.([]I) {
				if !predicate(item, i) {
					return false
				}
			}

			return true
		},
	}
}
