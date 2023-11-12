package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qahta0/stocksilo/config"
	"github.com/qahta0/stocksilo/repository"
	"github.com/qahta0/stocksilo/router"
)

func main() {
	// ##################### Init Application ##################### //
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// ##################### Middlewares ##################### //
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ##################### Start Configurations ##################### //
	config.LoadEnvs()
	config.DatabaseInit()

	// ##################### Migratations ##################### //
	repository.Migrate()

	// ##################### Test Database ##################### //
	gorm := config.DB()
	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}
	dbGorm.Ping()

	// ##################### Routes ##################### //
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Welcome to StockSilo API 🚀",
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "StockSilo API is healthy 🚀",
			"status":  "ok",
			"date":    time.Now().String(),
		})
	})

	router.StockRouter(e)

	// ##################### Start Server ##################### //
	e.Logger.Fatal(e.Start(":9090"))
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)
