package main

import (
	"profitCalculation/calc"
)

func main() {
	// TODO convert flag package for entering from CLI.
	// Please change the following three values.
	var investmentAmount float64 = 60000.00
	var purchasePrice float64 = 60.00
	var salePrice float64 = 64.04

	result := calc.New(investmentAmount, purchasePrice, salePrice)
	result.DisplayResult()
}
