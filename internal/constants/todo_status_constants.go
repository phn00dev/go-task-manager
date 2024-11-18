package constants

type TodoStatus string

const (
	New       TodoStatus = "new"
	Working   TodoStatus = "working"
	Done      TodoStatus = "done"
	Paused    TodoStatus = "paused"
	Cancelled TodoStatus = "cancelled"
)
