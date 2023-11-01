// service/stock.service.go

package service

import (
	"github.com/qahta0/stocksilo/model"
	"github.com/qahta0/stocksilo/repository"
)

func GetStocksService() ([]model.Stock, error) {
	return repository.GetStocks()
}

func CreateStockService(stock *model.Stock) (*model.Stock, error) {
	return repository.CreateStock(stock)
}

func GetStockService(id string) (*model.Stock, error) {
	return repository.GetStock(id)
}

func UpdateStockService(stock *model.Stock) error {
	return repository.UpdateStock(stock)
}

func DeleteStockService(id string) error {
	return repository.DeleteStock(id)
}
