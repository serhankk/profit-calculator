package utils

import (
	"fmt"
)


func FormatPercentageChangeBetweenTwoPrices(entryPrice float64, exitPrice float64, amount float64) string {

	prompt := fmt.Sprintf("Entry Price: %v\tExit Price: %v\tChange: %%%v", entryPrice, exitPrice, amount)
	return prompt
}


func FormatProfitByPrice(entryPrice float64, exitPrice float64, amount uint, profit float64) string {
	prompt := fmt.Sprintf("Entry Price: %v\tExit Price: %v\tProfit: %v", entryPrice, exitPrice, profit)
	return prompt
}


func FormatPriceAfterPercentageChange(entryPrice float64, percentageChange float64, newPrice float64) string {
	prompt := fmt.Sprintf("Entry Price: %v\tChange: %%%v\tNew Price: %v", entryPrice, percentageChange, newPrice)
	return prompt
}


func FormatTargetPriceByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPrice float64) string {
	prompt := fmt.Sprintf("Entry Price: %v\tAmount: %v\tTarget Profit: %v\tTarget Price: %v", entryPrice, amount, targetProfit, targetPrice)
	return prompt
}


func FormatTargetPercentageByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPercentage float64) string {
	prompt := fmt.Sprintf("Entry Price: %v\tAmount: %v\tTarget Profit: %v\tTarget Change: %v", entryPrice, amount, targetProfit, targetPercentage)
	return prompt
}
