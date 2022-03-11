package testdata

type Operation string

const (
	Union     Operation = "union"
	Connected Operation = "connected"
)

type TestStep struct {
	Op             Operation
	Args           [2]int
	ExpectedResult bool
}

type TestData struct {
	Name string
	N    int
	Ops  []TestStep
}
