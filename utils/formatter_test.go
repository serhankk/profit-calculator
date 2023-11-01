package utils

import (
	"regexp"
	"testing"
)

type formatChange struct {
	entryPrice, exitPrice, change float64
}

type formatProfit struct {
	entryPrice, exitPrice, amount, profit float64
}

var formatChangeData = [5]formatChange{
	{10, 20, 100},
	{27.5, 102.1, 271.27},
	{56, 22, -60.71},
	{999999, 555555, -44.44},
	{1, 1, 0},
}

var formatProfitData = [5]formatProfit{
	{1, 1, 2, 1},
	{10, 20, 22.36, 247.2},
	{220, 13, 160, -780},
	{155, 50, 155, 0},
	{155, 800, 150, -4000},
}

func TestFormatPercentageChangeBetweenTwoPrices(t *testing.T) {

	for _, test := range formatChangeData {

		result := FormatPercentageChangeBetweenTwoPrices(test.entryPrice, test.exitPrice, test.change, true)
		// "Entry Price: %v \t Exit Price: %v \t %s: %%%v"
		pattern := `^Entry Price: -?\d+(\.\d+)? \| Exit Price: -?\d+(\.\d+)? \| Change: %-?\d+(\.\d+)?$`
		t.Logf("pattern: %v\n", pattern)
		matched, err := regexp.MatchString(pattern, result)

		if err != nil {
			t.Log(err)
		}
		
		if matched {
			t.Logf("%v PASSED!", result)
		} else {
			t.Errorf("%v FAILED!", result)
		}
	}
}

func TestFormatProfitByPrice(t *testing.T) {

	for _, test := range formatProfitData {

		result := FormatProfitByPrice(test.entryPrice, test.exitPrice, uint(test.amount), test.profit, true)
		// "Entry Price: %v \t Exit Price: %v \t %s: %%%v"
		pattern := `^Entry Price: -?\d+(\.\d+)? \| Exit Price: -?\d+(\.\d+)? \| Profit: -?\d+(\.\d+)?$`
		t.Logf("pattern: %v\n", pattern)
		matched, err := regexp.MatchString(pattern, result)

		if err != nil {
			t.Log(err)
		}
		
		if matched {
			t.Logf("%v PASSED!", result)
		} else {
			t.Errorf("%v FAILED!", result)
		}
	}
}