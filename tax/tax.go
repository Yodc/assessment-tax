package tax

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

type TaxLevel struct {
	Level string      `json:"level"`
	Tax   json.Number `json:"tax"`
}

type Tax struct {
	Tax       json.Number `json:"tax"`
	TaxRefund json.Number `json:"taxRefund"`
	TaxLevel  []TaxLevel
}

type TaxCsv struct {
	TotalIncome json.Number `json:"totalIncome"`
	Tax         json.Number `json:"tax"`
	TaxRefund   json.Number `json:"taxRefund"`
}

type IncomeCsv struct {
	TotalIncome    float64 `json:"totalIncome"`
	Wht            float64 `json:"wht"`
	DonationAmount float64 `json:"donation"`
}

func TaxCalculationService(c echo.Context) error {

	var data Income

	err := c.Bind(&data)
	if err != nil {
		return err
	}

	if data.Wht < 0 {
		return c.String(http.StatusUnprocessableEntity, "Wht must positive number")
	}

	if data.Wht > data.TotalIncome {
		return c.String(http.StatusUnprocessableEntity, "Wht can't more than totalIncome")
	}

	for _, v := range data.Allowances {
		if v.Amount < 0 {
			return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("%s must positive number", v.AllowanceType))
		}
	}

	if data.TotalIncome <= 0 {
		return c.String(http.StatusUnprocessableEntity, "TotalIncome must positive number")
	}

	CalculationNetIncome(&data)
	taxAmount, taxLevel := TaxCalculation(data.NetIncome)
	taxAmountDeducted := DeductWht(&data, taxAmount)

	response := Tax{
		Tax:       toNumber(taxAmountDeducted),
		TaxRefund: toNumber(data.TaxRefund),
		TaxLevel:  taxLevel,
	}
	return c.JSON(http.StatusOK, response)
}

func ConvertDataFromCsv(data [][]string) ([]IncomeCsv, error) {

	var incomeData []IncomeCsv
	for idx, line := range data {
		if idx > 0 {

			if len(line) != 3 {
				return nil, errors.New(`data in record not equal 3`)
			}

			if strings.Trim(line[0], "") == "" || strings.Trim(line[1], "") == "" || strings.Trim(line[2], "") == "" {
				return nil, errors.New(`some data in record empty`)
			}

			totalIncome, err := strconv.ParseFloat(line[0], 64)
			if err != nil {
				return nil, errors.New(`totalIncome(line:%d) is not number`)
			}

			wht, err := strconv.ParseFloat(line[1], 64)
			if err != nil {
				return nil, errors.New(`wht(line:%d) is not number`)
			}

			donationAmount, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				return nil, errors.New(`donation(line:%d) is not number`)
			}

			item := IncomeCsv{
				TotalIncome:    totalIncome,
				Wht:            wht,
				DonationAmount: donationAmount,
			}

			incomeData = append(incomeData, item)
		}
	}

	return incomeData, nil
}

func TaxCalculationFromCSVService(c echo.Context) error {

	file, err := c.FormFile("taxFile")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}

	// Read CSV values using csv.Reader
	csvReader := csv.NewReader(src)
	data, err := csvReader.ReadAll()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	IncomeCsv, err := ConvertDataFromCsv(data)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	var taxs []TaxCsv

	for _, item := range IncomeCsv {
		income := Income{
			TotalIncome: item.TotalIncome,
			Wht:         item.Wht,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount:        item.DonationAmount,
				},
			},
		}

		CalculationNetIncome(&income)
		taxAmount, _ := TaxCalculation(income.NetIncome)
		taxAmountDeducted := DeductWht(&income, taxAmount)

		taxs = append(taxs, TaxCsv{
			TotalIncome: toNumber(income.TotalIncome),
			Tax:         toNumber(taxAmountDeducted),
			TaxRefund:   toNumber(income.TaxRefund),
		})

	}

	defer src.Close()
	return c.JSON(http.StatusOK, taxs)
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
