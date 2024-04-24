package tax

func TaxCalculation(netIncome float64) (float64, []TaxLevel) {
	taxBrackets := []struct {
		lower, upper, rate float64
	}{
		{0.00, 150000.00, 0.0},
		{150000.00, 500000.00, 0.1},
		{500000.00, 1000000.00, 0.15},
		{1000000.00, 2000000.00, 0.2},
		{2000000.00, 1e10, 0.35},
	}

	var arrResult []float64

	taxLevel := []TaxLevel{
		{
			Level: "0-150,000",
			Tax:   "0.00",
		},
		{
			Level: "150,001-500,000",
			Tax:   "0.00",
		},
		{
			Level: "500,001-1,000,000",
			Tax:   "0.00",
		},
		{
			Level: "1,000,001-2,000,000",
			Tax:   "0.00",
		},
		{
			Level: "2,000,001 ขึ้นไป",
			Tax:   "0.00",
		},
	}

	for idx, bracket := range taxBrackets {
		if netIncome > bracket.upper {
			arrResult = append(arrResult, (bracket.upper-bracket.lower)*bracket.rate)
			taxLevel[idx].Tax = toNumber(arrResult[idx])
		} else {
			arrResult = append(arrResult, (netIncome-bracket.lower)*bracket.rate)
			taxLevel[idx].Tax = toNumber(arrResult[idx])
			break
		}

	}

	tax := 0.00
	for _, result := range arrResult {
		tax += result
	}

	return tax, taxLevel
}
