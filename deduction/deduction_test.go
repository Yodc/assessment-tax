//go:build integation
// +build integation

package deduction_test

import (
	"testing"

	deduct "github.com/YodC/assessment-tax/deduction"
)

func TestDeduction(t *testing.T) {

	t.Run("Store or update personal deduction success when amount=60000", func(t *testing.T) {

		param := deduct.DeductionParam{
			Amount: 60000.00,
		}

		want := deduct.Deduction{
			DeductionAmount: 60000.00,
			DeductionType:   "personal_deduction",
		}

		got := deduct.InsertOrUpdatePersonalDeduction(param)

		if want.DeductionAmount != got.DeductionAmount && got.DeductionType == want.DeductionType && got.DeductionId > 0 {
			t.Errorf("InsertOrUpdatePersonalDeduction Failed")
		}
	})

	t.Run("Store or update k-receipt deduction success when amount=60000", func(t *testing.T) {

		param := deduct.DeductionParam{
			Amount: 60000.00,
		}

		want := deduct.Deduction{
			DeductionAmount: 60000.00,
			DeductionType:   "personal_deduction",
		}

		got := deduct.InsertOrUpdatePersonalDeduction(param)

		if want.DeductionAmount != got.DeductionAmount && got.DeductionType == want.DeductionType && got.DeductionId > 0 {
			t.Errorf("InsertOrUpdatePersonalDeduction Failed")
		}
	})

	t.Run("Store or update k-receipt deduction success when amount=50000", func(t *testing.T) {

		param := deduct.DeductionParam{
			Amount: 50000.00,
		}

		want := deduct.Deduction{
			DeductionAmount: 50000.00,
			DeductionType:   "k-receipt",
		}

		got := deduct.InsertOrUpdateKReceiptDeduction(param)

		if want.DeductionAmount != got.DeductionAmount && got.DeductionType == want.DeductionType && got.DeductionId > 0 {
			t.Errorf("InsertOrUpdatePersonalDeduction Failed")
		}
	})

	t.Run("Store or update donation deduction success when amount=100000", func(t *testing.T) {

		param := deduct.DeductionParam{
			Amount: 100000.00,
		}

		want := deduct.Deduction{
			DeductionAmount: 100000.00,
			DeductionType:   "donation",
		}

		got := deduct.InsertOrUpdateDonationDeduction(param)

		if want.DeductionAmount != got.DeductionAmount && got.DeductionType == want.DeductionType && got.DeductionId > 0 {
			t.Errorf("InsertOrUpdatePersonalDeduction Failed")
		}
	})
}
