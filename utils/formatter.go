package utils

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

func FormatPercentageChangeBetweenTwoPrices(entryPrice float64, exitPrice float64, change float64, colorEnabled bool) string {

	prompt := fmt.Sprintf("Entry Price: %v \tExit Price: %v \tChange: %%%v", entryPrice, exitPrice, change)
	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if change < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}
		fmt.Println(reflect.TypeOf(colorPrompt))
		prompt = fmt.Sprintf("Entry Price: %v \tExit Price: %v\t %s: %%%v", entryPrice, exitPrice, colorPrompt("Change"), change)
	}
	return prompt
}


func FormatProfitByPrice(entryPrice float64, exitPrice float64, amount uint, profit float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tExit Price: %v \tProfit: %v", entryPrice, exitPrice, profit)
	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if profit < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}

		prompt = fmt.Sprintf("Entry Price: %v \tExit Price: %v \t %s: %v", entryPrice, exitPrice, colorPrompt("Profit"), profit)
	}
	return prompt
}


func FormatPriceAfterPercentageChange(entryPrice float64, percentageChange float64, newPrice float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tChange: %%%v \tNew Price: %v", entryPrice, percentageChange, newPrice)
	if colorEnabled {
		
		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if newPrice < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}

		prompt = fmt.Sprintf("Entry Price: %v \tChange: %%%v \t %s: %v", entryPrice, percentageChange, colorPrompt("New Price"), newPrice)
	}
	return prompt
}


func FormatTargetPriceByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPrice float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \tTarget Price: %v", entryPrice, amount, targetProfit, targetPrice)
	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if targetPrice < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}

		prompt = fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \t %s: %v", entryPrice, amount, targetProfit, colorPrompt("Target Price"), targetPrice)
	}
	return prompt
}


func FormatTargetPercentageByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPercentage float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \tTarget Change: %v", entryPrice, amount, targetProfit, targetPercentage)
	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if targetPercentage < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}
		prompt = fmt.Sprintf("Entry Price: %v \tAmount: %v \tTarget Profit: %v \t %s: %v", entryPrice, amount, targetProfit, colorPrompt("Target Change"), targetPercentage)
	}
	return prompt
}
