package operators

type IOperator interface {
	GetOperatorSign() string
	ApplyOperation(left int, right int) float32
}
