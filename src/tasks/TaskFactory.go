package tasks

import (
	"../generators"
	"../operators"
)

type ITaskFactory interface {
	CreateTask() TaskToDo
}

type TaskFactory struct {
	numbersGenerator    generators.INumberGenerator
	operationsGenerator generators.IOperationGenerator
}

func (taskFactory *TaskFactory) CreateTask() TaskToDo {
	left := taskFactory.numbersGenerator.GetRandomNumber()
	right := taskFactory.numbersGenerator.GetRandomNumber()
	operator := taskFactory.operationsGenerator.GetRandomOperation()

	return TaskToDo{Left: left, Right: right, Operation: operator}
}

func CreateTaskFactory() TaskFactory {
	taskFactory := TaskFactory{
		numbersGenerator:    generators.CreateNumberGenerator(100),
		operationsGenerator: generators.CreateOperationGenerator(),
	}

	addOperator := operators.AddOperator{}
	minusOperator := operators.MinusOperator{}

	taskFactory.operationsGenerator.AddOperation(addOperator)
	taskFactory.operationsGenerator.AddOperation(addOperator)
	taskFactory.operationsGenerator.AddOperation(minusOperator)

	return taskFactory
}
