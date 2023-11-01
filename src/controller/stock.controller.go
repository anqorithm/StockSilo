// controller/stock.controller.go

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qahta0/stocksilo/model"
	"github.com/qahta0/stocksilo/service"
)

func GetStocks(c echo.Context) error {
	stocks, err := service.GetStocksService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stocks)
}

func CreateStock(c echo.Context) error {
	var stock model.Stock
	if err := c.Bind(&stock); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newStock, err := service.CreateStockService(&stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, newStock)
}

func GetStock(c echo.Context) error {
	id := c.Param("id")
	stock, err := service.GetStockService(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stock)
}

func UpdateStock(c echo.Context) error {
	id := c.Param("id")
	var stock model.Stock
	if err := c.Bind(&stock); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	stock.ID = id
	if err := service.UpdateStockService(&stock); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, stock)
}

func DeleteStock(c echo.Context) error {
	id := c.Param("id")
	if err := service.DeleteStockService(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Stock deleted successfully"})
}
