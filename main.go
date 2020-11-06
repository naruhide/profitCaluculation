package main

import (
	"fmt"

	"profitCalculation/calc"
)

func main() {
	// TODO convert flag package for entering from CLI.
	// Please change the following three values.
	var investmentAmount float64 = 37936.59
	var purchasePrice float64 = 51.97
	var salePrice float64 = 58.55

	result := calc.New(investmentAmount, purchasePrice, salePrice)
	fmt.Println(result.NumberOfSharesPurchased)
	fmt.Println(result.RateOfUp)
	fmt.Println(result.ProfitAfterTax)
	fmt.Println(result.TaxAndFee)
	fmt.Println(result.TotalAssets)
}
