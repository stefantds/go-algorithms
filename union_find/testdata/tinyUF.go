package testdata

var TinyUF = TestData{
	Name: "Tiny test data",
	N:    10,
	Ops: []TestStep{
		{
			Op:             Connected,
			Args:           [2]int{4, 3},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{4, 3},
		},
		{
			Op:             Connected,
			Args:           [2]int{3, 8},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{3, 8},
		},
		{
			Op:             Connected,
			Args:           [2]int{6, 5},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{6, 5},
		},
		{
			Op:             Connected,
			Args:           [2]int{9, 4},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{9, 4},
		},
		{
			Op:             Connected,
			Args:           [2]int{2, 1},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{2, 1},
		},
		{
			Op:             Connected,
			Args:           [2]int{8, 9},
			ExpectedResult: true,
		},
		{
			Op:   Union,
			Args: [2]int{8, 9},
		},
		{
			Op:             Connected,
			Args:           [2]int{5, 0},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{5, 0},
		},
		{
			Op:             Connected,
			Args:           [2]int{7, 2},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{7, 2},
		},
		{
			Op:             Connected,
			Args:           [2]int{6, 1},
			ExpectedResult: false,
		},
		{
			Op:   Union,
			Args: [2]int{6, 1},
		},
		{
			Op:             Connected,
			Args:           [2]int{1, 0},
			ExpectedResult: true,
		},
		{
			Op:   Union,
			Args: [2]int{1, 0},
		},
		{
			Op:             Connected,
			Args:           [2]int{6, 7},
			ExpectedResult: true,
		},
		{
			Op:   Union,
			Args: [2]int{6, 7},
		},
		{
			Op:             Connected,
			Args:           [2]int{6, 7},
			ExpectedResult: true,
		},
	},
}
