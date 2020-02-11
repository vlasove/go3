package main

import "testing"

type CheckerData struct {
	message  string
	expected int
}

var dataCheck = []CheckerData{
	{"opopopopoooooop", 6},
	{"ooo", 3},
	{"ppppppppppopppppp", 1},
	{"ooooooo", 7},
	{"ppppppppppppppp", 0},
	{"ppooppooppooppooo", 3},
	{"ooopopopopopopop", 3},
	{"ppppppppppppppppop", 1},
	{"ppppppppppppppppppppppppppppppppooop", 3},
	{"pop", 1},
}

func TestTableChecker(t *testing.T) {
	for _, test := range dataCheck {
		if out := Checker(test.message); out != test.expected {
			t.Error("error happend --- answer: {}, checker: {}", test.expected, out)
		}
	}
}
