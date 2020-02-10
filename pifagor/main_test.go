package main

import "testing"

type PifagoreanData struct {
	A        int
	B        int
	C        int
	expected string
}

var data = []PifagoreanData{
	{3, 4, 5, "Yes"},
	{4, 3, 5, "Yes"},
	{1, 2, 3, "No"},
	{3, 3, 3, "No"},
	{8, 10, 6, "Yes"},
}

func TestTablePifagor(t *testing.T) {
	for _, test := range data {
		if out := isPifagor(test.A, test.B, test.C); out != test.expected {
			t.Error("error happend during testing")
		}
	}
}
