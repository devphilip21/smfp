package smfp

type (
	Pipeline[I any] struct {
		workers []*Worker
	}
)

func (p *Pipeline[I]) Execute(input I) any {
	if len(p.workers) == 0 {
		panic(ErrEmptyPipe)
	}

	var data any = input
	for _, worker := range p.workers {
		data = worker.Execute(data)
	}

	return data
}

func Pipe[I any](workers []*Worker) *Pipeline[I] {
	return &Pipeline[I]{
		workers: workers,
	}
}
