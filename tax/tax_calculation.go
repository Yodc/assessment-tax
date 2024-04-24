package tax

func TaxCalculation(netIncome float64) float64 {
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

	for _, bracket := range taxBrackets {
		if netIncome > bracket.upper {
			arrResult = append(arrResult, (bracket.upper-bracket.lower)*bracket.rate)
		} else {
			arrResult = append(arrResult, (netIncome-bracket.lower)*bracket.rate)
			break
		}

	}

	tax := 0.00
	for _, result := range arrResult {
		tax += result
	}

	return tax
}
