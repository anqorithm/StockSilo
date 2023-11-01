package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/qahta0/stocksilo/config"
	"github.com/qahta0/stocksilo/repository"
	"github.com/qahta0/stocksilo/router"
)

func main() {
	e := echo.New()
	config.LoadEnvs()
	config.DatabaseInit()
	repository.Migrate()

	gorm := config.DB()
	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}
	dbGorm.Ping()

	// ##################### Start Routes ##################### //

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

	// ##################### End Routes ##################### //

	e.Logger.Fatal(e.Start(":9090"))
}
