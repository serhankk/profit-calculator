package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"calc.com/profit-calculator/utils"
)
var isColorEnabled bool = true

func main(){
	options := []string{"change-percentage", "profit-by-price", "price-after-change", "price-for-profit", "change-for-profit"}
	if !isOptionPassed() {
		fmt.Println("Options are: ")
		for _, val := range options {
			fmt.Printf("-> %v\n", val)
		}
		os.Exit(0)
	}

	option := flag.String("option", options[0], fmt.Sprintf("%v", options))
	entryPrice := flag.Float64("buy", 0, "Entry price")
	exitPrice := flag.Float64("sell", 0, "Exit price")
	amount := flag.Uint("amount", 0, "Amount")
	targetProfit := flag.Float64("profit", 0, "Target profit")
	changePercentage := flag.Float64("change", 0, "Change percentage")

	flag.Parse()

	isOptionValid := slices.Contains(options, *option)
	if !isOptionValid {
		fmt.Printf("invalid option: %v\n", *option)
		fmt.Printf("options: %v\n", options)
	}


	switch *option {
	case options[0]:
		fmt.Println(calculateChangePercentageResult(*entryPrice, *exitPrice))
	case options[1]:
		fmt.Println(calculateProfitByPrice(*entryPrice, *exitPrice, *amount))
	case options[2]:
		fmt.Println(calculatePriceAfterPercentageChange(*entryPrice, *changePercentage))
	case options[3]:
		fmt.Println(calculateTargetPriceByTargetProfit(*entryPrice, *amount, *targetProfit))
	case options[4]:
		fmt.Println(calculatePercentageByTargetProfit(*entryPrice, *amount, *targetProfit))
	}
}

func isOptionPassed() (bool) {
		return len(os.Args) > 1
}

func calculateChangePercentageResult(entryPrice float64, exitPrice float64) string {
	changePercentageResult := utils.CalculatePercentageChangeBetweenTwoPrices(entryPrice, exitPrice)
	formattedChangePercentageResult := utils.FormatPercentageChangeBetweenTwoPrices(entryPrice, exitPrice, changePercentageResult, isColorEnabled)

	return formattedChangePercentageResult
}

func calculateProfitByPrice(entryPrice float64, exitPrice float64, amount uint) string {
	profitByPriceResult := utils.CalculateProfitByPrice(entryPrice, exitPrice, uint(amount))
	formattedProfitByPriceResult := utils.FormatProfitByPrice(entryPrice, exitPrice, uint(amount), profitByPriceResult, isColorEnabled)

	return formattedProfitByPriceResult
}

func calculatePriceAfterPercentageChange(entryPrice float64, changePercetage float64) string {
	priceAfterPercetageChange := utils.CalculatePriceAfterPercentageChange(entryPrice, float64(changePercetage))
	formattedPriceAfterPercetageChange := utils.FormatPriceAfterPercentageChange(entryPrice, float64(changePercetage), priceAfterPercetageChange, isColorEnabled)

	return formattedPriceAfterPercetageChange
}

func calculateTargetPriceByTargetProfit(entryPrice float64, amount uint, targetProfit float64) string {
	targetPriceByTargetProfit := utils.CalculateTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	formattedTargetPriceByTargetProfit :=  utils.FormatTargetPriceByTargetProfit(entryPrice, uint(amount), float64(targetProfit), targetPriceByTargetProfit, isColorEnabled)

	return formattedTargetPriceByTargetProfit
}

func calculatePercentageByTargetProfit(entryPrice float64, amount uint, targetProfit float64) string {
	targetPercentageByTargetProfit := utils.CalculateTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit))
	formattedTargetPercentageByTargetProfit := utils.FormatTargetPercentageByTargetProfit(entryPrice, uint(amount), float64(targetProfit), targetPercentageByTargetProfit, isColorEnabled)

	return formattedTargetPercentageByTargetProfit
}