package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qahta0/stocksilo/model"
	"github.com/qahta0/stocksilo/service"
	"github.com/qahta0/stocksilo/utils"
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
	var createStockRequest model.CreateStockRequest

	if err := c.Bind(&createStockRequest); err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	if err := c.Validate(&createStockRequest); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, "Validation failed for stock data", err)
	}

	newStock, err := service.CreateStockService(&createStockRequest)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, "Error creating stock", err)
	}

	return c.JSON(http.StatusCreated, newStock)
}

func UpdateStock(c echo.Context) error {
	var updateStockRequest model.UpdateStockRequest
	if err := c.Bind(&updateStockRequest); err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}
	if err := c.Validate(&updateStockRequest); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, "Validation failed for stock data", err)
	}
	stockID := c.Param("id")
	updatedStock, err := service.UpdateStockService(stockID, &updateStockRequest)
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, "Error updating stock", err)
	}
	return c.JSON(http.StatusOK, updatedStock)
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
