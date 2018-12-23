package tasks

import (
	"../operators"
)

type TaskToDo struct {
	Left   int
	Right  int
	Result float32

	Operation operators.IOperator
}
