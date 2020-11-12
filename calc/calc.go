// calc package includes calculation of investment result simulation.
package calc

import "fmt"

const (
	domesticFeesPerTransaction               float64 = 0.45      // Represented by percentage.
	maxDomesticFees                          float64 = 20        // Dollar.
	localTransactionCostsPerOneContractPrice float64 = 0.0000221 // Only at the time of sale.
	transferProfitTaxation                   float64 = 0.20315   // Represented by percentage.
	//exchangeRates                            float64 = 104.5     // Yen and dollar.
)

// InvestResult represents calculation result of investment.
type InvestResult struct {
	NumberOfSharesPurchased float64 `json:"number_of_shares_purchased"`
	RateOfUp                float64 `json:"rate_of_up"`
	ProfitAfterTax          float64 `json:"profit_after_tax"`
	TaxAndFee               float64 `json:"tax_and_fee"`
	TotalAssets             float64 `json:"total_assets"`
	AssetGrowthRate         float64 `json:"asset_growth_rate"`
}

// DisplayResult displays members of InvestResult struct.
func (result *InvestResult) DisplayResult() {
	fmt.Printf("NumberOfSharesPurchased: %.2f\n", result.NumberOfSharesPurchased)
	fmt.Printf("RateOfUp: %.2f\n", result.RateOfUp)
	fmt.Printf("ProfitAfterTax: %.2f\n", result.ProfitAfterTax)
	fmt.Printf("TaxAndFee: %.2f\n", result.TaxAndFee)
	fmt.Printf("TotalAssets: %.2f\n", result.TotalAssets)
}

// New initializes InvestResult struct.
func New(investmentAmount, purchasePrice, salePrice float64) *InvestResult {
	numberOfSharesPurchased := investmentAmount / purchasePrice
	rateOfUp := salePrice / purchasePrice

	valuation := rateOfUp * investmentAmount
	profitBeforeTax := valuation - investmentAmount

	domesticFees := 2 * calcDomesticFees(investmentAmount)
	fee := domesticFees + (localTransactionCostsPerOneContractPrice * numberOfSharesPurchased)
	tax := profitBeforeTax * transferProfitTaxation
	taxAndFee := tax + fee
	profitAfterTax := profitBeforeTax - taxAndFee

	totalAssets := investmentAmount + profitAfterTax
	assetGrowthRate := totalAssets / investmentAmount

	result := &InvestResult{
		NumberOfSharesPurchased: numberOfSharesPurchased,
		RateOfUp:                rateOfUp,
		ProfitAfterTax:          profitBeforeTax,
		TaxAndFee:               taxAndFee,
		TotalAssets:             totalAssets,
		AssetGrowthRate:         assetGrowthRate,
	}
	return result
}

// calcDomesticFees calclates domestic fees. Up to 20$.
func calcDomesticFees(investmentAmount float64) float64 {
	domesticFees := investmentAmount * domesticFeesPerTransaction
	if domesticFees < 20 {
		return domesticFees
	} else {
		return maxDomesticFees
	}
}

//func Truncatefloat(num float64) float64 {

//}

//func ExchangeDollarToYen(dollar float64) float64 {
//	return dollar * exchangeRates
//}
