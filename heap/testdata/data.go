package testdata

type Operation string

const (
	Insert Operation = "insert"
	Max    Operation = "max"
	DelMax Operation = "delMax"
)

type TestStep struct {
	Op             Operation
	Arg            int
	ExpectedResult int
}

type TestData struct {
	Name string
	Ops  []TestStep
}
