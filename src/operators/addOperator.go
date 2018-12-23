package operators

type AddOperator struct {
}

func (addOperator AddOperator) ApplyOperation(left int, right int) float32 {
	return float32(left + right)
}

func (addOperator AddOperator) GetOperatorSign() string {
	return "+"
}
