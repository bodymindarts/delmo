package delmo

type Tasks map[string]Task

type Task struct{}

func NewTasks(configs []TaskConfig) Tasks {
	tasks := Tasks{}
	return tasks
}
