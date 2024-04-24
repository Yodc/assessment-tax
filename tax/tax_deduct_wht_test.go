package tax

import (
	"testing"
)

func TestDeductWht(t *testing.T) {

	t.Run("Should return taxAmount=0, taxRefund=0 when taxAmount=0 and wht=0", func(t *testing.T) {

		data := Income{
			Wht: 0.00,
		}
		taxAmount := 0.00

		want_1 := 0.00
		want_2 := 0.00

		got_1 := DeductWht(&data, taxAmount)
		get_2 := data.TaxRefund

		if got_1 != want_1 && data.TaxRefund == want_2 {
			t.Errorf("TestNetIncome(%f, %f) = %f, %f; want %f, %f", taxAmount, data.Wht, got_1, get_2, want_1, want_2)
		}
	})

	t.Run("Should return taxAmount=0, taxRefund=50 when taxAmount=0 and wht=50", func(t *testing.T) {

		data := Income{
			Wht: 50.00,
		}
		taxAmount := 0.00

		want_1 := 0.00
		want_2 := 50.00

		got_1 := DeductWht(&data, taxAmount)
		get_2 := data.TaxRefund

		if got_1 != want_1 && data.TaxRefund == want_2 {
			t.Errorf("TestNetIncome(%f, %f) = %f, %f; want %f, %f", taxAmount, data.Wht, got_1, get_2, want_1, want_2)
		}
	})

	t.Run("Should return taxAmount=50, taxRefund=0 when taxAmount=100 and wht=50", func(t *testing.T) {

		data := Income{
			Wht: 50.00,
		}
		taxAmount := 100.00

		want_1 := 50.00
		want_2 := 0.00

		got_1 := DeductWht(&data, taxAmount)
		get_2 := data.TaxRefund

		if got_1 != want_1 && data.TaxRefund == want_2 {
			t.Errorf("TestNetIncome(%f, %f) = %f, %f; want %f, %f", taxAmount, data.Wht, got_1, get_2, want_1, want_2)
		}
	})

	t.Run("Should return taxAmount=100, taxRefund=0 when taxAmount=100 and wht=0", func(t *testing.T) {

		data := Income{
			Wht: 0.00,
		}
		taxAmount := 100.00

		want_1 := 100.00
		want_2 := 0.00

		got_1 := DeductWht(&data, taxAmount)
		get_2 := data.TaxRefund

		if got_1 != want_1 && data.TaxRefund == want_2 {
			t.Errorf("TestNetIncome(%f, %f) = %f, %f; want %f, %f", taxAmount, data.Wht, got_1, get_2, want_1, want_2)
		}
	})

}
