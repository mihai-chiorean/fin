package calculators

// ApplyCompountInterest calculates the result after years of compounded interest
func AppyCompoundInterest(sum, years, rate float64) float64 {
	result := sum
	for i := 1.00; i <= years; i++ {
		result = result + (rate * result)
	}

	return result
}
