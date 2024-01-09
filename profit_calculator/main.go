package main

import "fmt"

func main() {
	revenue := getInput("Revenue: ")
	expenses := getInput("Expenses ")
	taxRate := getInput("Tax Rate: ")

	ebt, eat, ratio := calculateEarnings(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.0f\n", ebt)
	fmt.Printf("Earnings After Tax: %.0f\n", eat)
	fmt.Printf("Earning Ratio: %.2f\n", ratio)
}

func getInput(text string) (input float64) {
	fmt.Print(text)
	fmt.Scan(&input)
	return input
}

func calculateEarnings(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat
	return ebt, eat, ratio
}
