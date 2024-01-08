package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("Earnings Before Tax: %.0f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.0f\n", earningsAfterTax)
	fmt.Printf("Earning Ratio: %.0f\n", ratio)
}
