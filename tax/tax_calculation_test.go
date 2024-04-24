package tax_test

import (
	"testing"

	tax "github.com/YodC/assessment-tax/tax"
)

func TestTaxCalculation(t *testing.T) {

	t.Run("Should return tax=0 when total_income=0", func(t *testing.T) {
		total_income := 0.00

		want := 0.00

		got, _ := tax.TaxCalculation(total_income)

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=0 when total_income=150000", func(t *testing.T) {
		total_income := 150000.00

		got, _ := tax.TaxCalculation(total_income)

		want := 0.00

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=0.1 when total_income=150001", func(t *testing.T) {
		total_income := 150001.00

		got, _ := tax.TaxCalculation(total_income)

		want := 0.10

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=50000 when total_income=500000", func(t *testing.T) {
		total_income := 500000.00

		got, _ := tax.TaxCalculation(total_income)

		want := 35000.00

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=50150 when total_income=500001", func(t *testing.T) {
		total_income := 500001.00

		got, _ := tax.TaxCalculation(total_income)

		want := 35000.00 + 0.15

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=110000 when total_income=1000000", func(t *testing.T) {
		total_income := 1000000.00

		got, _ := tax.TaxCalculation(total_income)

		want := 35000.00 + 75000.00

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=125200 when total_income=1000001", func(t *testing.T) {
		total_income := 1000001.00

		got, _ := tax.TaxCalculation(total_income)

		want := 35000.00 + 75000.00 + 0.20

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=100 when total_income=2000000", func(t *testing.T) {
		total_income := 2000000.00

		got, _ := tax.TaxCalculation(total_income)

		want := 35000.00 + 75000.00 + 200000.00

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})

	t.Run("Should return tax=100 when total_income=2000001", func(t *testing.T) {
		total_income := 2000001.00

		got, _ := tax.TaxCalculation(total_income)

		want := 35000.00 + 75000.00 + 200000.00 + 0.35

		if got != want {
			t.Errorf("TaxCalculation(%f) = %f; want %f", total_income, got, want)
		}
	})
}
