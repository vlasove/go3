package main

import "testing"

type PyramidTester struct {
	h        int
	expected string
}

var dataPyramid = []PyramidTester{
	{1, "@"},
	{2, " @\n@@@"},
	{3, "  @\n @@@\n@@@@@"},
	{4, "   @\n  @@@\n @@@@@\n@@@@@@@"},
	{5, "    @\n   @@@\n  @@@@@\n @@@@@@@\n@@@@@@@@@"},
	{6, "     @\n    @@@\n   @@@@@\n  @@@@@@@\n @@@@@@@@@\n@@@@@@@@@@@"},
}

func TestTablePyramid(t *testing.T) {
	for _, test := range dataPyramid {
		if out := Pyramid(test.h); out != test.expected {
			t.Error("error happend during testing")
		}
	}
}
