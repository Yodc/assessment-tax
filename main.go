package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	deduct "github.com/YodC/assessment-tax/deduction"
	tax "github.com/YodC/assessment-tax/tax"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	godotenv.Load(".env")
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	// nomal
	e.POST("/tax/calculations", tax.TaxCalculationService)
	e.POST("/tax/calculations/upload-csv", tax.TaxCalculationFromCSVService)

	// admin group
	g := e.Group("admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Check if username and password match your desired credentials
		if username == os.Getenv("ADMIN_USERNAME") && password == os.Getenv("ADMIN_PASSWORD") {
			return true, nil
		}
		return false, nil
	}))
	g.POST("/deductions/personal", deduct.InsertOrUpdatePersonalDeductionService)
	g.POST("/deductions/k-receipt", deduct.InsertOrUpdateKReceiptDeductionService)
	g.POST("/deductions/donation", deduct.InsertOrUpdateDonationDeductionService)

	go func() {
		if err := e.Start(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("shutting down the server")

}
