package main

import (
	"testing"
)

func TestToSnafu(t *testing.T) {
	type testCase struct {
		decimal int
		snafu   string
	}

	testCases := []testCase{
		{decimal: 1, snafu: "1"},
		{decimal: 2, snafu: "2"},
		{decimal: 3, snafu: "1="},
		{decimal: 4, snafu: "1-"},
		{decimal: 5, snafu: "10"},
		{decimal: 6, snafu: "11"},
		{decimal: 7, snafu: "12"},
		{decimal: 8, snafu: "2="},
		{decimal: 9, snafu: "2-"},
		{decimal: 10, snafu: "20"},
		{decimal: 15, snafu: "1=0"},
		{decimal: 20, snafu: "1-0"},
		{decimal: 2022, snafu: "1=11-2"},
		{decimal: 12345, snafu: "1-0---0"},
		{decimal: 314159265, snafu: "1121-1110-1=0"},
	}

	for _, testCase := range testCases {
		snafu := toSnafu(testCase.decimal, 15, "")
		if snafu != testCase.snafu {
			t.Errorf("toSnafu conversion failed for %d: %s vs %s", testCase.decimal, snafu, testCase.snafu)
		}

		decimal := toDecimal(testCase.snafu)
		if decimal != testCase.decimal {
			t.Errorf("toDecimal conversion failed for %d: %s vs %s", testCase.decimal, snafu, testCase.snafu)
		}
	}
}
