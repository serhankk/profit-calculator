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

type calculatePrice struct {
	entryPrice, change, expected float64
}

type calculateTargetPrice struct {
	entryPrice, amount, profit, expected float64
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

var calculatePriceData = [5]calculatePrice{
	{200, 4, 208},
	{55.8, 20, 66.96},
	{82, -8, 75.44},
	{1, 1, 1.01},
	{50.1, -1, 49.59},
}
var calculateTargetPriceData = [5]calculateTargetPrice{
	{200, 10, 208, 220.8},
	{8, 1, 20, 28},
	{7.65, 150, 300, 9.65},
	{30, 100, -150, 28.5},
	{1, 1, 1, 2},
}

var calculateTargetChangeData = [5]calculateTargetPrice{
	{200, 10, 208, 10.4},
	{8, 1, 20, 250},
	{7.65, 150, 300, 26.14},
	{30, 100, -150, -5},
	{1, 1, 1, 100},
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

func TestCalculatePriceAfterPercentageChange(t *testing.T) {
	for num, test := range calculatePriceData {
		result := CalculatePriceAfterPercentageChange(test.entryPrice, test.change)

		if result != test.expected {
			t.Errorf("%v. CalculatePriceAfterPercentageChange(%v, %v) FAILED! Expected: %v | Got: %v", num, test.entryPrice, test.change, test.expected, result)
		} else {
			t.Logf("%v. CalculatePriceAfterPercentageChange(%v, %v) PASSED! Expected: %v | Got: %v", num, test.entryPrice, test.change, test.expected, result)
		}
	}
}

func TestCalculateTargetPriceByTargetProfit(t *testing.T){
	for num, test := range calculateTargetPriceData {
		result := CalculateTargetPriceByTargetProfit(test.entryPrice, uint(test.amount), test.profit)

		if result != test.expected {
			t.Errorf("%v. CalculateTargetPriceByTargetProfit(%v, %v, %v) FAILED! Expected: %v | Got: %v", num, test.entryPrice, test.amount, test.profit, test.expected, result)
		} else {
			t.Logf("%v. CalculateTargetPriceByTargetProfit(%v, %v, %v) PASSED! Expected: %v | Got: %v", num, test.entryPrice, test.amount, test.profit, test.expected, result)
		}
	}
}

func TestCalculateTargetPercentageByTargetProfit(t *testing.T){
	for num, test := range calculateTargetChangeData {
		result := CalculateTargetPercentageByTargetProfit(test.entryPrice, uint(test.amount), test.profit)

		if result != test.expected {
			t.Errorf("%v. CalculateTargetPriceByTargetProfit(%v, %v, %v) FAILED! Expected: %v | Got: %v", num, test.entryPrice, test.amount, test.profit, test.expected, result)
		} else {
			t.Logf("%v. CalculateTargetPriceByTargetProfit(%v, %v, %v) PASSED! Expected: %v | Got: %v", num, test.entryPrice, test.amount, test.profit, test.expected, result)
		}
	}
}