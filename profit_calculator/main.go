package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue := getInput("Revenue: ")
	expenses := getInput("Expenses ")
	taxRate := getInput("Tax Rate: ")

	ebt, eat, ratio, err := calculateEarnings(revenue, expenses, taxRate)

	if err != nil {
		fmt.Println(err)
		return
	}

	saveResultsToFile(ebt, eat, ratio)
}

func getInput(text string) (input float64) {
	fmt.Print(text)
	fmt.Scan(&input)
	return input
}

func calculateEarnings(revenue, expenses, taxRate float64) (ebt float64, eat float64, ratio float64, err error) {
	if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		return ebt, eat, ratio, errors.New("Invalid input. All values must be greater than 0.")
	}

	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat

	return ebt, eat, ratio, err
}

func saveResultsToFile(ebt, eat, ratio float64) {
	stringToOutput := fmt.Sprintf("EBT: %.0f\nEAT: %.0f\nRatio: %.2f", ebt, eat, ratio)
	os.WriteFile("results.txt", []byte(stringToOutput), 0644)
	fmt.Println("Results saved to results.txt")
}
