package utils

import "testing"

type addTest struct {
	entryPrice, exitPrice, expected float64
}

var calculatePercentageChangeData = [5]addTest{
	{10, 20, 100},
	{27.5, 102.1, 271.27},
	{56, 22, -60.71},
	{999999999, 555555555, -44.44},
	{1, 1, 0},
}

func TestCalculatePercentageChangeBetweenTwoPrices(t *testing.T) {

	for _, test := range calculatePercentageChangeData {
		result := CalculatePercentageChangeBetweenTwoPrices(test.entryPrice, test.exitPrice)

		if result != test.expected {
			t.Errorf("CalculatePercentageChangeBetweenTwoPrices(%f, %f) FAILED! Expected: %v | Got: %v", test.entryPrice, test.exitPrice, test.expected, result)
		} else {
			t.Logf("CalculatePercentageChangeBetweenTwoPrices(%f, %f) PASSED! Expected: %v | Got: %v", test.entryPrice, test.exitPrice, test.expected, result)
		}
	}
}

func TestCalculateProfitByPrice(t *testing.T) {
	result := CalculateProfitByPrice(10, 20, 10)

	if result != 100 {
		t.Errorf("CalculateProfitByPrice(10, 20, 10) FAILED! Expected: %v | Got: %v", 100, result)
	} else {
		t.Logf("CalculateProfitByPrice(10, 20, 10) PASSED! Expected: %v | Got: %v", 100, result)
	}
}
