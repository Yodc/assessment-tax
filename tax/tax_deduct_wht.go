package tax

func DeductWht(data *Income, taxAmount float64) float64 {

	if data.Wht > taxAmount {
		data.TaxRefund = data.Wht - taxAmount
		return 0
	} else {
		data.TaxRefund = 0
		return taxAmount - data.Wht
	}

}
