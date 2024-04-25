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

func InsertOrUpdatePersonalDeductionService(c echo.Context) error {
	db := database.InitDBPostgres()

	var param DeductionParam
	err := c.Bind(&param)
	if err != nil {
		return err
	}

	deduction := Deduction{
		DeductionType:   "personal_deduction",
		DeductionAmount: param.Amount,
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "deduction_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"deduction_amount"}),
	}).Create(&deduction)

	return c.JSON(http.StatusOK, PersonalDeductionResponse{PersonalDeduction: json.Number(toNumber(deduction.DeductionAmount))})
}

func InsertOrUpdateKReceiptDeductionService(c echo.Context) error {
	db := database.InitDBPostgres()

	var param DeductionParam
	err := c.Bind(&param)
	if err != nil {
		return err
	}

	deduction := Deduction{
		DeductionType:   "k-receipt",
		DeductionAmount: param.Amount,
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "deduction_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"deduction_amount"}),
	}).Create(&deduction)

	return c.JSON(http.StatusOK, PersonalDeductionResponse{PersonalDeduction: json.Number(toNumber(deduction.DeductionAmount))})
}

func InsertOrUpdateDonationDeductionService(c echo.Context) error {
	db := database.InitDBPostgres()

	var param DeductionParam
	err := c.Bind(&param)
	if err != nil {
		return err
	}

	deduction := Deduction{
		DeductionType:   "donation",
		DeductionAmount: param.Amount,
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "deduction_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"deduction_amount"}),
	}).Create(&deduction)

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
