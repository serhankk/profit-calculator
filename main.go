package main

import (
	"calc.com/profit-calculator/utils"
)


func main(){

	entryPrice := 100.2
	exitPrice := 121.5
	amount := 25

	changePercetange := 25
	targetProfit := 155
	
	changePercentageResult := utils.CalculatePercentageChangeBetweenTwoPrices(entryPrice, exitPrice)
	formattedChangePercentageResult := utils.FormatPercentageChangeBetweenTwoPrices(entryPrice, exitPrice, changePercentageResult)

	profitByPriceResult := utils.CalculateProfitByPrice(entryPrice, exitPrice, uint(amount))
	formattedProfitByPriceResult := utils.FormatProfitByPrice(entryPrice, exitPrice, uint(amount), profitByPriceResult)
	
	priceAfterPercetangeChange := utils.CalculatePriceAfterPercentageChange(entryPrice, float64(changePercetange))
	formattedPriceAfterPercetangeChange := utils.FormatPriceAfterPercentageChange(entryPrice, float64(changePercetange), priceAfterPercetangeChange)

	targetPriceByTargetProfit := utils.CalculateTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	formattedTargetPriceByTargetProfit :=  utils.FormatTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit), targetPriceByTargetProfit)

	targetPercentageByTargetProfit := utils.CalculateTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	formattedTargetPercentageByTargetProfit := utils.FormatTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit), targetPercentageByTargetProfit)


}
