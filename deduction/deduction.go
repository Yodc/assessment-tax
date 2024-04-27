package deduction

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/YodC/assessment-tax/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type Deduction struct {
	DeductionId     int     `json:"deduction_id" gorm:"primaryKey"`
	DeductionType   string  `json:"deduction_type"`
	DeductionAmount float64 `json:"deduction_amount"`
}

type DeductionParam struct {
	Amount float64 `json:"amount"`
}

type PersonalDeductionResponse struct {
	PersonalDeduction json.Number `json:"personalDeduction"`
}

func GetDeductionByDeductionType(deductionType string) Deduction {
	db := database.InitDBPostgres()
	var deduction Deduction
	db.First(&deduction, "deduction_type = ?", deductionType)

	return deduction
}

func InsertOrUpdatePersonalDeduction(param DeductionParam) Deduction {
	db := database.InitDBPostgres()

	deduction := Deduction{
		DeductionType:   "personal_deduction",
		DeductionAmount: param.Amount,
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "deduction_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"deduction_amount"}),
	}).Create(&deduction)

	return deduction
}

func InsertOrUpdatePersonalDeductionService(c echo.Context) error {

	var param DeductionParam
	err := c.Bind(&param)
	if err != nil {
		return err
	}

	if param.Amount < 10000 || param.Amount > 100000 {
		return c.String(http.StatusUnprocessableEntity, "personal deduction is number between 10000 - 100000")
	}

	deduction := InsertOrUpdatePersonalDeduction(param)

	return c.JSON(http.StatusOK, PersonalDeductionResponse{PersonalDeduction: json.Number(toNumber(deduction.DeductionAmount))})
}

func InsertOrUpdateKReceiptDeduction(param DeductionParam) Deduction {
	db := database.InitDBPostgres()

	deduction := Deduction{
		DeductionType:   "k-receipt",
		DeductionAmount: param.Amount,
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "deduction_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"deduction_amount"}),
	}).Create(&deduction)

	return deduction
}

func InsertOrUpdateKReceiptDeductionService(c echo.Context) error {

	var param DeductionParam
	err := c.Bind(&param)
	if err != nil {
		return err
	}

	if param.Amount < 0 || param.Amount > 100000 {
		return c.String(http.StatusUnprocessableEntity, "k-receipt is number between 0 - 100000")
	}

	deduction := InsertOrUpdateKReceiptDeduction(param)
	return c.JSON(http.StatusOK, PersonalDeductionResponse{PersonalDeduction: json.Number(toNumber(deduction.DeductionAmount))})
}

func InsertOrUpdateDonationDeduction(param DeductionParam) Deduction {
	db := database.InitDBPostgres()

	deduction := Deduction{
		DeductionType:   "donation",
		DeductionAmount: param.Amount,
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "deduction_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"deduction_amount"}),
	}).Create(&deduction)

	return deduction
}

func InsertOrUpdateDonationDeductionService(c echo.Context) error {

	var param DeductionParam
	err := c.Bind(&param)
	if err != nil {
		return err
	}

	deduction := InsertOrUpdateDonationDeduction(param)

	return c.JSON(http.StatusOK, PersonalDeductionResponse{PersonalDeduction: json.Number(toNumber(deduction.DeductionAmount))})
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
