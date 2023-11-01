package router

import (
	"github.com/labstack/echo/v4"
	"github.com/qahta0/stocksilo/controller"
)

func StockRouter(e *echo.Echo) {
	p := e.Group("/api/v1")
	p.GET("/stocks", controller.GetStocks)
	p.POST("/stocks", controller.CreateStock)
	p.GET("/stocks/:id", controller.GetStock)
	p.PUT("/stocks/:id", controller.UpdateStock)
	p.DELETE("/stocks/:id", controller.DeleteStock)
}
