package tax

import "fmt"

func CalculationNetIncome(data *Income) {

	deductAmount := 0.00

	deductAmount += 60000

	for _, item := range data.Allowances {
		if item.AllowanceType == "donation" {
			if item.Amount > 100000.00 {
				deductAmount += 100000.00
			} else {
				deductAmount += item.Amount

			}
		}
	}

	fmt.Println(deductAmount)
	netIncome := data.TotalIncome - deductAmount
	if netIncome <= 0 {
		data.NetIncome = 0
	} else {
		data.NetIncome = netIncome
	}

}
