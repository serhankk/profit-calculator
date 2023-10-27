package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func FormatPercentageChangeBetweenTwoPrices(entryPrice float64, exitPrice float64, change float64, colorEnabled bool) string {

	prompt := fmt.Sprintf("Entry Price: %v \tExit Price: %v \tChange: %%%v", entryPrice, exitPrice, change)
	if colorEnabled {
		green := color.New(color.FgGreen).SprintFunc()
		prompt = fmt.Sprintf("Entry Price: %v \tExit Price: %v\t %s: %%%v", entryPrice, exitPrice, green("Change"), change)
	}
	return prompt
}


func FormatProfitByPrice(entryPrice float64, exitPrice float64, amount uint, profit float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tExit Price: %v \tProfit: %v", entryPrice, exitPrice, profit)
	if colorEnabled {
		green := color.New(color.FgGreen).SprintFunc()
		prompt = fmt.Sprintf("Entry Price: %v \tExit Price: %v \t %s: %v", entryPrice, exitPrice, green("Profit"), profit)
	}
	return prompt
}


func FormatPriceAfterPercentageChange(entryPrice float64, percentageChange float64, newPrice float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tChange: %%%v \tNew Price: %v", entryPrice, percentageChange, newPrice)
	if colorEnabled {
		green := color.New(color.FgGreen).SprintFunc()
		prompt = fmt.Sprintf("Entry Price: %v \tChange: %%%v \t %s: %v", entryPrice, percentageChange, green("New Price"), newPrice)
	}
	return prompt
}


func FormatTargetPriceByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPrice float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \tTarget Price: %v", entryPrice, amount, targetProfit, targetPrice)
	if colorEnabled {
		green := color.New(color.FgGreen).SprintFunc()
		prompt = fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \t %s: %v", entryPrice, amount, targetProfit, green("Target Price"), targetPrice)
	}
	return prompt
}


func FormatTargetPercentageByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPercentage float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \tTarget Change: %v", entryPrice, amount, targetProfit, targetPercentage)
	if colorEnabled {
		green := color.New(color.FgGreen).SprintFunc()
		prompt = fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \t %s: %v", entryPrice, amount, targetProfit, green("Target Change"), targetPercentage)
	}
	return prompt
}
