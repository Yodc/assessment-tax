package tax

import deduct "github.com/YodC/assessment-tax/deduction"

func CalculationNetIncome(data *Income) {

	deductAmount := 0.00

	deductPersonalAmount := deduct.GetDeductionByDeductionType("personal_deduction").DeductionAmount
	deductKReceiptAmount := deduct.GetDeductionByDeductionType("k-receipt").DeductionAmount
	deductDonationAmount := deduct.GetDeductionByDeductionType("donation").DeductionAmount

	deductAmount += deductPersonalAmount

	for _, item := range data.Allowances {
		if item.AllowanceType == "donation" {
			if item.Amount > deductDonationAmount {
				deductAmount += deductDonationAmount
			} else {
				deductAmount += item.Amount

			}
		}

		if item.AllowanceType == "k-receipt" {
			if item.Amount > deductKReceiptAmount {
				deductAmount += deductKReceiptAmount
			} else {
				deductAmount += item.Amount
			}
		}
	}

	netIncome := data.TotalIncome - deductAmount
	if netIncome <= 0 {
		data.NetIncome = 0
	} else {
		data.NetIncome = netIncome
	}

}
