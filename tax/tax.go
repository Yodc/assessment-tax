package tax

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Income struct {
	TotalIncome float64     `json:"totalIncome"`
	Wht         float64     `json:"wht"`
	NetIncome   float64     `json:"netIncome"`
	Allowances  []Allowance `json:"allowances"`
	TaxRefund   float64
}

type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type Tax struct {
	Tax       json.Number `json:"tax"`
	TaxRefund json.Number `json:"taxRefund"`
}

func TaxCalculationService(c echo.Context) error {

	var data Income

	err := c.Bind(&data)
	if err != nil {
		return err
	}

	CalculationNetIncome(&data)
	taxAmount := TaxCalculation(data.NetIncome)
	taxAmountDeducted := DeductWht(&data, taxAmount)

	response := Tax{
		Tax:       toNumber(taxAmountDeducted),
		TaxRefund: toNumber(data.TaxRefund),
	}
	return c.JSON(http.StatusOK, response)
}

func toNumber(f float64) json.Number {
	var s string
	if f == float64(int64(f)) {
		s = fmt.Sprintf("%.1f", f) // 1 decimal if integer
	} else {
		s = fmt.Sprint(f)
	}
	return json.Number(s)
}
