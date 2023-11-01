package utils

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

func floatLimiter(inputNumber float64) float64 {
	limitedFloat := math.Floor(inputNumber * 100) / 100
	fmt.Println(limitedFloat)
	return limitedFloat
}

func FormatPercentageChangeBetweenTwoPrices(entryPrice float64, exitPrice float64, change float64, colorEnabled bool) string {

	entryPrice = floatLimiter(entryPrice)
	exitPrice = floatLimiter(exitPrice)
	change = floatLimiter(change)

	prompt := fmt.Sprintf("Entry Price: %v | Exit Price: %v | Change: %%%v", entryPrice, exitPrice, change)


	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if change < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}
		prompt = fmt.Sprintf("Entry Price: %v | Exit Price: %v | %s: %%%v", entryPrice, exitPrice, colorPrompt("Change"), change)
	}
	return prompt
}


func FormatProfitByPrice(entryPrice float64, exitPrice float64, amount uint, profit float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v | Exit Price: %v | Profit: %v", entryPrice, exitPrice, profit)
	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if profit < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}

		prompt = fmt.Sprintf("Entry Price: %v | Exit Price: %v | %s: %v", entryPrice, exitPrice, colorPrompt("Profit"), profit)
	}
	return prompt
}


func FormatPriceAfterPercentageChange(entryPrice float64, percentageChange float64, newPrice float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v | Change: %%%v | New Price: %v", entryPrice, percentageChange, newPrice)
	if colorEnabled {
		
		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if newPrice < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}

		prompt = fmt.Sprintf("Entry Price: %v | Change: %%%v | %s: %v", entryPrice, percentageChange, colorPrompt("New Price"), newPrice)
	}
	return prompt
}


func FormatTargetPriceByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPrice float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v | Amount: %v | Target Profit: %v | Target Price: %v", entryPrice, amount, targetProfit, targetPrice)
	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if targetPrice < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}

		prompt = fmt.Sprintf("Entry Price: %v | Amount: %v | Target Profit: %v | %s: %v", entryPrice, amount, targetProfit, colorPrompt("Target Price"), targetPrice)
	}
	return prompt
}


func FormatTargetPercentageByTargetProfit(entryPrice float64, amount uint, targetProfit float64, targetPercentage float64, colorEnabled bool) string {
	prompt := fmt.Sprintf("Entry Price: %v | Amount: %v | Target Profit: %v | Target Change: %v", entryPrice, amount, targetProfit, targetPercentage)
	if colorEnabled {

		colorPrompt := color.New(color.FgGreen).SprintFunc()
		if targetPercentage < 0 {
		colorPrompt = color.New(color.FgRed).SprintFunc()
	}
		prompt = fmt.Sprintf("Entry Price: %v | Amount: %v | Target Profit: %v | %s: %v", entryPrice, amount, targetProfit, colorPrompt("Target Change"), targetPercentage)
	}
	return prompt
}
