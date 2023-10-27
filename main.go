package main

import (
	"fmt"

	"calc.com/profit-calculator/utils"
)


func main(){

	entryPrice := 100.2
	exitPrice := 121.5
	amount := 25

	changePercetange := 25
	targetProfit := 155
	
	changePercentageResult := utils.CalculatePercentageChangeBetweenTwoPrices(entryPrice, exitPrice)
	fmt.Println(utils.FormatPercentageChangeBetweenTwoPrices(entryPrice, exitPrice, changePercentageResult))

	myVar2 := utils.CalculateProfitByPrice(entryPrice, exitPrice, uint(amount))
	// fmt.Println(myVar2)
	fmt.Println(utils.FormatProfitByPrice(entryPrice, exitPrice, uint(amount), myVar2))
	
	myVar3 := utils.CalculatePriceAfterPercentageChange(entryPrice, float64(changePercetange))
	// fmt.Println(myVar3)
	fmt.Println(utils.FormatPriceAfterPercentageChange(entryPrice, float64(changePercetange), myVar3))

	myVar4 := utils.CalculateTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	// fmt.Println(myVar4)
	fmt.Println(utils.FormatTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit), myVar4))

	myVar5 := utils.CalculateTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	// fmt.Println(myVar5)
	fmt.Println(utils.FormatTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit), myVar5))


}
