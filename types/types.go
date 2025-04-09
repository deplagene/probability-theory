package types

type ProbabilityData struct {
	Title		string
	Theory      string
	Formula     string
	ExampleText string
	Example     Example
}

type Example struct {
	N int
	M int
	P int
}

type ProbabilityService interface {
	Calculate(n, m int) (float64, error)
}