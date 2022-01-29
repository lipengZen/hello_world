package main

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	testSlice := []struct{ a, b, c int }{
		{a: 3, b: 4, c: 5},
		{a: 6, b: 8, c: 1},
		{a: 10, b: 12, c: 13},
		{a: 30000, b: 40000, c: 50000},
	}

	for _, testCase := range testSlice {
		if c := calcTriangle(testCase.a, testCase.b); c != testCase.c {
			// t.Logf("calc triangle err a=%d, b=%d, c=%d but get c:%d", testCase.a, testCase.b, testCase.c, c)
			t.Errorf("calc triangle err a=%d, b=%d, c=%d but get c:%d", testCase.a, testCase.b, testCase.c, c)
		}
	}
}
