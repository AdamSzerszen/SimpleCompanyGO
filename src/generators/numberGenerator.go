package generators

import (
	"math/rand"
	"time"
)

type INumberGenerator interface {
	GetRandomNumber() int
}

type NumberGenerator struct {
	numbersRange int
}

func (numberGenerator NumberGenerator) GetRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % numberGenerator.numbersRange
}

func CreateNumberGenerator(numbersRande int) NumberGenerator {
	generator := NumberGenerator{numbersRange: numbersRande}

	return generator
}
