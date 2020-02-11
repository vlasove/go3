package main

import "testing"

type SoupTest struct {
	days     int
	expected string
}

var dataSoup = []SoupTest{
	{1, "щи"},
	{2, "щи\nборщ"},
	{3, "щи\nборщ\nщавелевый суп"},
	{8, "щи\nборщ\nщавелевый суп\nовсяный суп\nсуп по-чабански\nщи\nборщ\nщавелевый суп"},
	{5, "щи\nборщ\nщавелевый суп\nовсяный суп\nсуп по-чабански"},
	{10, "щи\nборщ\nщавелевый суп\nовсяный суп\nсуп по-чабански\nщи\nборщ\nщавелевый суп\nовсяный суп\nсуп по-чабански"},
}

func TestTableSoupMaker(t *testing.T) {
	for _, test := range dataSoup {
		if out := SoupMaker(test.days); out != test.expected {
			t.Error("error happend during testing in day {}", test.days)
		}
	}
}
