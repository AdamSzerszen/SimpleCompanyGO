package generators

import (
	"../operators"
	"math/rand"
	"time"
)

type IOperationGenerator interface {
	GetRandomOperation() *operators.IOperator
	AddOperation(operator *operators.IOperator)
}

type OperationGenerator struct {
	availableOperations []*operators.IOperator
}

func (operationGenerator OperationGenerator) GetRandomOperation() *operators.IOperator {
	rand.Seed(time.Now().UnixNano())
	operationsRange := len(operationGenerator.availableOperations)
	randIndex := rand.Int() % operationsRange
	return operationGenerator.availableOperations[randIndex]
}

func (operationGenerator OperationGenerator) AddOperation(operator *operators.IOperator) {
	operationGenerator.availableOperations = append(operationGenerator.availableOperations, operator)
}

func CreateOperationGenerator() *IOperationGenerator {
	generator := new(OperationGenerator{availableOperations: []*operators.IOperator{}})

	return generator
}
