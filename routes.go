package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Finances",
		})
	})
	// INVESTMENT SIMULATION
	app.Get("/investments", func(c *fiber.Ctx) error {
		return c.Render("investments", fiber.Map{
			"title": "Investments",
		})
	})
	app.Post("/investments", func(c *fiber.Ctx) error {
		i := new(InvestmentRequest)
		if err := c.BodyParser(i); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		total, e := simulateInvestment(i.BaseValue, i.DurationMonths, i.MonthlyRate/100.0, i.ApplyTax)
		if e != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": e,
			})
		}
		response := InvestmentResponse{
			BaseValue:      i.BaseValue,
			MonthlyRate:    i.MonthlyRate,
			DurationMonths: i.DurationMonths,
			ApplyTax:       i.ApplyTax,
			TotalValue:     total,
		}
		return c.Render("partials/investments_response", fiber.Map{
			"BaseValue":      fmt.Sprintf("%.2f", response.BaseValue),
			"MonthlyRate":    fmt.Sprintf("%.2f", response.MonthlyRate),
			"DurationMonths": response.DurationMonths,
			"ApplyTax":       response.ApplyTax,
			"TotalValue":     fmt.Sprintf("%.2f", response.TotalValue),
			"Earnings":       fmt.Sprintf("%.2f", response.TotalValue-response.BaseValue),
			"Profit":         fmt.Sprintf("%.2f", (response.TotalValue/response.BaseValue)*100.0-100.0),
		})
	})

	// CDB COMPARISON
	app.Get("/cdb", func(c *fiber.Ctx) error {
		return c.Render("cdb", fiber.Map{
			"title": "CDB Comparison",
		})
	})
	app.Put("/cdb", func(c *fiber.Ctx) error {
		r := new(CDBCompRequest)
		if err := c.BodyParser(r); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		trueCDBRate, err := realCdbRate(r.BaseCDBRate, r.DurationMonths)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.Render("partials/cdb_result", fiber.Map{
			"TrueCDBRate": fmt.Sprintf("%.2f", trueCDBRate),
			"DiffPct":     fmt.Sprintf("%.2f", trueCDBRate/r.LCIRate),
		})
	})

	// SHOPPING INFORMATION
	app.Get("/shopping", func(c *fiber.Ctx) error {
		return c.Render("shopping", fiber.Map{
			"title": "Shopping Info",
		})
	})
}
