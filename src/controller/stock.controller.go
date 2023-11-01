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
		c.Logger().Errorf("Error retrieving stocks: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}
	if len(stocks) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"success": "false",
			"message": "No stocks found",
		})
	}
	return c.JSON(http.StatusOK, stocks)
}

func CreateStock(c echo.Context) error {
	var stock model.Stock
	if err := c.Bind(&stock); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}
	if err := c.Validate(&stock); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Validation failed for stock data",
			"details": err.Error(),
		})
	}
	newStock, err := service.CreateStockService(&stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}
	return c.JSON(http.StatusCreated, newStock)
}

func GetStock(c echo.Context) error {
	id := c.Param("id")
	stock, err := service.GetStockService(id)
	if err != nil {
		c.Logger().Errorf("Error retrieving stock with ID %s: %v", id, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}
	if stock == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Stock not found"})
	}
	return c.JSON(http.StatusOK, stock)
}

func UpdateStock(c echo.Context) error {
	id := c.Param("id")
	var stock model.Stock
	if err := c.Bind(&stock); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}
	stock.ID = id
	existingStock, err := service.GetStockService(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}
	if existingStock == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Stock not found",
		})
	}
	if err := service.UpdateStockService(&stock); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Failed to update stock",
			"details": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, stock)
}

func DeleteStock(c echo.Context) error {
	id := c.Param("id")
	_, err := service.GetStockService(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}
	if err := service.DeleteStockService(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting stock"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Stock deleted successfully"})
}
