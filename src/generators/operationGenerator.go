package generators

import (
	"math/rand"
	"time"
)

type IOperationGenerator interface {
	GetRandomOperation() string
}

type OperationGenerator struct {
	availableOperations []string
}

func (operationGenerator OperationGenerator) GetRandomOperation() string {
	rand.Seed(time.Now().UnixNano())
	operationsRange := len(operationGenerator.availableOperations)
	randIndex := rand.Int() % operationsRange
	return operationGenerator.availableOperations[randIndex]
}

func CreateOperationGenerator(availableOperations []string) *IOperationGenerator {
	generator := new(OperationGenerator{availableOperations: availableOperations})

	return generator
}
