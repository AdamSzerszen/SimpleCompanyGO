package operators

type MinusOperator struct {
}

func (minusOperator MinusOperator) ApplyOperation(left int, right int) float32 {
	return float32(left - right)
}

func (minusOperator MinusOperator) GetOperatorSign() string {
	return "-"
}
