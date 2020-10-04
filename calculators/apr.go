package calculators

import (
	"math"
)

// APRInfo describes information about a calculated apr
type APRInfo struct {
	MonthlyPayment float64
	Total          float64
	TotalInterest  float64
}

// round2dec rounds to the nearest 2 decimal
func round2dec(n float64) float64 {
	return math.Round(n*100) / 100
}

// APR - returns data about the APR values
//	principal = the amount borrowed
//  payments = number of monthly payments e.g. 30 years = 360
//  rate = the base percentage rate of the loan. A 5.25% Annual Rate should be passed in as 0.0525 NOT 5.25
//  costs = the loan closing costs e.g. origination fee, broker fees, etc.
func APR(principal, payments, rate, costs float64) APRInfo {
	rate = rate / 12.00
	mp := round2dec(((principal + costs) * rate * math.Pow(1.00+rate, payments)) / (math.Pow(1.00+rate, payments) - 1))

	testrate := rate
	i := 1
	testdiff := testrate
	for i <= 100 {
		testres := ((testrate * math.Pow(1.00+testrate, payments)) / (math.Pow(1.00+testrate, payments) - 1)) - (mp / principal)

		if testres < 0.0000001 {
			break
		}
		if testres < 0 {
			testrate += testdiff
		} else {
			testdiff /= 2
		}
		i++
	}
	testrate *= 12

	return APRInfo{
		round2dec(mp),
		round2dec(mp * payments),
		round2dec((mp * payments) - principal - costs),
	}
}
