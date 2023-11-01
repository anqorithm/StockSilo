package router

import (
	"github.com/labstack/echo/v4"
	"github.com/qahta0/stocksilo/controller"
)

func StockRouter(e *echo.Echo) {
	e.GET("/stocks", controller.GetStocks)
	e.POST("/stocks", controller.CreateStock)
	e.GET("/stocks/:id", controller.GetStock)
	e.PUT("/stocks/:id", controller.UpdateStock)
	e.DELETE("/stocks/:id", controller.DeleteStock)
}
