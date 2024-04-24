package tax

func CalculationNetIncome(data *Income) {

	deductAmount := 0.00

	deductAmount += 60000

	if data.TotalIncome <= 60000 {
		data.NetIncome = 0
	} else {
		data.NetIncome = data.TotalIncome - deductAmount
	}

}
