package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"calc.com/profit-calculator/utils"
)


func main(){

	/*
	entryPrice := 100.2
	exitPrice := 121.5
	
	
	amount := 25

	changePercetage := 25
	targetProfit := 155

	*/

	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	option = strings.Replace(option, "\n", "", -1)
	if err != nil {
		fmt.Println(err)
	}

	intro_msg := `Select one of the options (number):
	
1. Calculate change percentage by ENTRY PRICE and EXIT PRICE

2. Calculate profit by ENTRY PRICE, EXIT PRICE and AMOUNT

3. Calculate price after percentage change by ENTRY PRICE and CHANGE PERCENTAGE

4. Calculate target price by TARGET PROFIT

5. Calculate target percentage change by TARGET PROFIT.`

	fmt.Println(intro_msg)

	switch option {
	case "1":
		fmt.Printf("sectin: %v\n", option)
	case "2":
		fmt.Println("sectin: xcvx")
	}

	
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