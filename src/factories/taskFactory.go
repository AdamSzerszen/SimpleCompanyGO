package factories

import (
	"../generators"
	"../operators"
	"../tasks"
)

type ITaskFactory interface {
	CreateTask() tasks.TaskToDo
}

type TaskFactory struct {
	NumbersGenerator    *generators.INumberGenerator
	OperationsGenerator *generators.IOperationGenerator
}

func (taskFactory TaskFactory) CreateTask() tasks.TaskToDo {
	left := (*taskFactory.NumbersGenerator).GetRandomNumber()
	right := (*taskFactory.NumbersGenerator).GetRandomNumber()
	operator := (*taskFactory.OperationsGenerator).GetRandomOperation()

	return tasks.TaskToDo{Left: left, Right: right, Operation: operator}
}

func CreateTaskFactory() ITaskFactory {

	taskFactory := TaskFactory{
		NumbersGenerator:    generators.CreateNumberGenerator(100),
		OperationsGenerator: generators.CreateOperationGenerator(),
	}

	addOperator := new(operators.AddOperator{})
	minusOperator := new(operators.MinusOperator{})

	(*taskFactory.OperationsGenerator).AddOperation(addOperator)
	(*taskFactory.OperationsGenerator).AddOperation(minusOperator)

	return taskFactory
}
