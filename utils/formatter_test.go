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

type formatPrice struct {
	entryPrice, change, newPrice float64
}

type formatTargetPrice struct {
	entryPrice, amount, profit, targetPrice float64
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

var formatPriceData = [5]formatPrice{
	{200, 4, 208},
	{55.8, 20, 66.96},
	{82, -8, 75.44},
	{1, 1, 1.01},
	{50.1, -1, 49.59},
}

var formatTargetPriceData = [5]formatTargetPrice{
	{200, 10, 208, 220.8},
	{8, 1, 20, 28},
	{7.65, 150, 300, 9.65},
	{30, 100, -150, 28.5},
	{1, 1, 1, 2},
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

func TestFormatPriceAfterPercentageChange(t *testing.T) {

	for _, test := range formatPriceData {

		result := FormatPriceAfterPercentageChange(test.entryPrice, test.change, test.newPrice, true)
		// "Entry Price: %v \t Exit Price: %v \t %s: %%%v"
		pattern := `^Entry Price: -?\d+(\.\d+)? \| Change: %-?\d+(\.\d+)? \| New Price: -?\d+(\.\d+)?$`
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

func TestFormatTargetPriceByTargetProfit(t *testing.T) {

	for _, test := range formatTargetPriceData {

		result := FormatTargetPriceByTargetProfit(test.entryPrice, uint(test.amount), test.profit, test.targetPrice, true)
		// "Entry Price: %v \t Exit Price: %v \t %s: %%%v"
		pattern := `^Entry Price: -?\d+(\.\d+)? \| Amount: -?\d+(\.\d+)? \| Target Profit: -?\d+(\.\d+)? \| Target Price: -?\d+(\.\d+)?$`
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