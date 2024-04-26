//go:build integation
// +build integation

package tax

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCalculationNetIncome(t *testing.T) {

	godotenv.Load("../.env")
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

	t.Run("Should return netIncome=1 when totalIncome=60001", func(t *testing.T) {

		data := Income{
			TotalIncome: 60001.00,
		}

		want := 1.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=900 when totalIncome=61000, donation=100", func(t *testing.T) {

		data := Income{
			TotalIncome: 61000.00,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount:        100.00,
				},
			},
		}

		want := 900.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=0 when totalIncome=61000, donation=1100", func(t *testing.T) {

		data := Income{
			TotalIncome: 61000.00,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount:        1100.00,
				},
			},
		}

		want := 0.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=900 when totalIncome=61000, donation=1000", func(t *testing.T) {

		data := Income{
			TotalIncome: 61000.00,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount:        1000.00,
				},
			},
		}

		want := 0.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=900 when totalIncome=61000, k-receipt=100", func(t *testing.T) {

		data := Income{
			TotalIncome: 61000.00,
			Allowances: []Allowance{
				{
					AllowanceType: "k-receipt",
					Amount:        100.00,
				},
			},
		}

		want := 900.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=0 when totalIncome=61000, k-receipt=1100", func(t *testing.T) {

		data := Income{
			TotalIncome: 61000.00,
			Allowances: []Allowance{
				{
					AllowanceType: "k-receipt",
					Amount:        1100.00,
				},
			},
		}

		want := 0.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

	t.Run("Should return netIncome=900 when totalIncome=61000, k-receipt=1000", func(t *testing.T) {

		data := Income{
			TotalIncome: 61000.00,
			Allowances: []Allowance{
				{
					AllowanceType: "k-receipt",
					Amount:        1000.00,
				},
			},
		}

		want := 0.00

		CalculationNetIncome(&data)

		if data.NetIncome != want {
			t.Errorf("TestNetIncome(%f) = %f; want %f", data.TotalIncome, data.NetIncome, want)
		}
	})

}
