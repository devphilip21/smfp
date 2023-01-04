package smfp

type (
	Task func(input any) (output any)

	Worker struct {
		task Task
	}
)

func (w *Worker) Execute(input any) any {
	output := w.task(input)

	return output
}
