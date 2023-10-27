package main

import (
	"fmt"

	"calc.com/profit-calculator/utils"
)


func main(){

	entryPrice := 100.2
	exitPrice := 121.5
	
	/*
	amount := 25

	changePercetage := 25
	targetProfit := 155

	*/

	fmt.Println(calculateChangePercentageResult(entryPrice, exitPrice))
	
}

func calculateChangePercentageResult(entryPrice float64, exitPrice float64) string {
	changePercentageResult := utils.CalculatePercentageChangeBetweenTwoPrices(entryPrice, exitPrice)
	formattedChangePercentageResult := utils.FormatPercentageChangeBetweenTwoPrices(entryPrice, exitPrice, changePercentageResult)

	return formattedChangePercentageResult
}

func calculateProfitByPrice(entryPrice float64, exitPrice float64, amount uint) string {
	profitByPriceResult := utils.CalculateProfitByPrice(entryPrice, exitPrice, uint(amount))
	formattedProfitByPriceResult := utils.FormatProfitByPrice(entryPrice, exitPrice, uint(amount), profitByPriceResult)

	return formattedProfitByPriceResult
}

func calculatePriceAfterPercentageChange(entryPrice float64, changePercetage float64) string {
	priceAfterPercetageChange := utils.CalculatePriceAfterPercentageChange(entryPrice, float64(changePercetage))
	formattedPriceAfterPercetageChange := utils.FormatPriceAfterPercentageChange(entryPrice, float64(changePercetage), priceAfterPercetageChange)

	return formattedPriceAfterPercetageChange
}

func calculateTargetPriceByTargetProfit(entryPrice float64, amount uint, targetProfit float64) string {
	targetPriceByTargetProfit := utils.CalculateTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	formattedTargetPriceByTargetProfit :=  utils.FormatTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit), targetPriceByTargetProfit)

	return formattedTargetPriceByTargetProfit
}

func calculatePercentageByTargetProfit(entryPrice float64, amount uint, targetProfit float64) string {
	targetPercentageByTargetProfit := utils.CalculateTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	formattedTargetPercentageByTargetProfit := utils.FormatTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit), targetPercentageByTargetProfit)

	return formattedTargetPercentageByTargetProfit
}