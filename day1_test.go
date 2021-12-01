package main

import "testing"

func TestDay1Example(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	const expected = 7

	actual := numberOfTimesDepthIncreases(input)

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
