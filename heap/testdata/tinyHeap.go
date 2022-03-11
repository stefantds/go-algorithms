package testdata

var TinyHeapData = TestData{
	Name: "Tiny test data",
	Ops: []TestStep{
		{
			Op:  Insert,
			Arg: 4,
		},
		{
			Op:  Insert,
			Arg: 2,
		},
		{
			Op:             Max,
			ExpectedResult: 4,
		},
		{
			Op:  Insert,
			Arg: 5,
		},
		{
			Op:             Max,
			ExpectedResult: 5,
		},
		{
			Op:             DelMax,
			ExpectedResult: 5,
		},
		{
			Op:             Max,
			ExpectedResult: 4,
		},
		{
			Op:             DelMax,
			ExpectedResult: 4,
		},
		{
			Op:             Max,
			ExpectedResult: 2,
		},
		{
			Op:  Insert,
			Arg: 3,
		},
		{
			Op:             Max,
			ExpectedResult: 3,
		},
	},
}
