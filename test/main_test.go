package main

import "testing"

type LinearTest struct {
	B        int
	C        int
	expected string
}

var linearData = []LinearTest{
	{2, 2, "Has 1 root"},
	{4, 0, "Has 1 root"},
	{0, 2, "Has No roots"},
}

func TestTableLinear(t *testing.T) {
	for _, test := range linearData {
		if out := Linear(test.B, test.C); out != test.expected {
			t.Error("error happend during linear test")
		}
	}
}

type QuadraticTest struct {
	A        int
	B        int
	C        int
	expected string
}

var quadraticData = []QuadraticTest{
	{2, 1, 2, "Has No roots"},
	{1, -3, 2, "Has 2 roots"},
	{1, -200, 2, "Has 2 roots"},
}

func TestTableQuadratic(t *testing.T) {
	for _, test := range quadraticData {
		if out := Quadratic(test.A, test.B, test.C); out != test.expected {
			t.Error("error happend during quadratic test")
		}
	}
}

type SolveTest struct {
	A        int
	B        int
	C        int
	expected string
}

var solveData = []SolveTest{
	{0, 2, 2, "Has 1 root"},
	{1, -3, 2, "Has 2 roots"},
	{0, 0, 2, "Has No roots"},
}

func TestTableSolve(t *testing.T) {
	for _, test := range solveData {
		if out := Solve(test.A, test.B, test.C); out != test.expected {
			t.Error("error happend during solve test")
		}
	}
}
