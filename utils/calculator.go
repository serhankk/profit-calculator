package utils

import (
	"fmt"
	"os"
	"strconv"
)


func CalculatePercentageChangeBetweenTwoPrices(entryPrice float64, exitPrice float64) float64{
	stringResult := fmt.Sprintf("%.2f", (exitPrice - entryPrice) / entryPrice * 100)
	result, err := strconv.ParseFloat(stringResult, 64)
	if err != nil {
		os.Exit(1)
	} 
		
	return result	
}

func CalculateProfitByPrice(entryPrice float64, exitPrice float64, amount uint) float64 {
	result := (exitPrice - entryPrice) * float64(amount)
	return result
}

func CalculatePriceAfterPercentageChange(entryPrice float64, percentageChange float64) float64 {
	result := entryPrice + (entryPrice * percentageChange / 100)
	return result
}

func CalculateTargetPriceByTargetProfit(entryPrice float64, amount uint, targetProfit float64) float64 {
	exitPrice := (targetProfit / float64(amount)) + entryPrice
	return exitPrice
}

func CalculateTargetPercentageByTargetProfit(entryPrice float64, amount uint, targetProfit float64) float64 {
	targetPercentage := (targetProfit / (entryPrice * float64(amount)) * 100)
	return targetPercentage
}