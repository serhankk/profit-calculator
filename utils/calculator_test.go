package utils

import (
	"testing"
)

type calculateChange struct {
	entryPrice, exitPrice, expected float64
}

type calculateProfit struct {
	entryPrice, amount, exitPrice, expected float64
}

var calculateChangeData = [5]calculateChange{
	{10, 20, 100},
	{27.5, 102.1, 271.27},
	{56, 22, -60.71},
	{999999999, 555555555, -44.44},
	{1, 1, 0},
}

var calculateProfitData = [5]calculateProfit{
	{1, 1, 2, 1},
	{10, 20, 22.36, 247.2},
	{220, 13, 160, -780},
	{155, 50, 155, 0},
	{155, 800, 150, -4000},
}

func TestCalculatePercentageChangeBetweenTwoPrices(t *testing.T) {

	for num, test := range calculateChangeData {
		result := CalculatePercentageChangeBetweenTwoPrices(test.entryPrice, test.exitPrice)

		if result != test.expected {
			t.Errorf("%v. CalculatePercentageChangeBetweenTwoPrices(%f, %f) FAILED! Expected: %v | Got: %v", num, test.entryPrice, test.exitPrice, test.expected, result)
		} else {
			t.Logf("%v. CalculatePercentageChangeBetweenTwoPrices(%f, %f) PASSED! Expected: %v | Got: %v", num, test.entryPrice, test.exitPrice, test.expected, result)
		}
	}
}

func TestCalculateProfitByPrice(t *testing.T) {

	for num, test := range calculateProfitData {
		result := CalculateProfitByPrice(test.entryPrice, test.exitPrice, uint(test.amount))

		if result != test.expected {
			t.Errorf("%v. CalculateProfitByPrice(%v, %v, %v) FAILED! Expected: %v | Got: %v", num, test.entryPrice, test.amount, test.exitPrice, test.expected, result)
		} else {
			t.Logf("%v. CalculateProfitByPrice(%v, %v, %v) PASSED! Expected: %v | Got: %v", num, test.entryPrice, test.amount, test.exitPrice, test.expected, result)
		}
	}
}
