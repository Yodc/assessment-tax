package tax

import (
	"testing"
)

func TestCalculationNetIncome(t *testing.T) {

	t.Run("Should return netIncome=0 when totalIncome=60000", func(t *testing.T) {

		data := Income{
			TotalIncome: 60000.00,
		}

		want := 0.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=0 when totalIncome=0", func(t *testing.T) {

		data := Income{
			TotalIncome: 60000.00,
		}

		want := 0.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=0 when totalIncome=60001", func(t *testing.T) {

		data := Income{
			TotalIncome: 60001.00,
		}

		want := 1.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

}
