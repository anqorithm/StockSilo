package service

import (
	"reflect"

	"github.com/qahta0/stocksilo/model"
	"github.com/qahta0/stocksilo/repository"
)

func GetStocksService() ([]model.Stock, error) {
	return repository.GetStocks()
}

func CreateStockService(createReq *model.CreateStockRequest) (*model.Stock, error) {
	stock := model.Stock{
		Ticker:       createReq.Ticker,
		Name:         createReq.Name,
		Open:         createReq.Open,
		Close:        createReq.Close,
		High:         createReq.High,
		Low:          createReq.Low,
		Volume:       createReq.Volume,
		Date:         createReq.Date,
		CurrentPrice: createReq.CurrentPrice,
	}
	newStock, err := repository.CreateStock(&stock)
	if err != nil {
		return nil, err
	}

	return newStock, nil
}

func GetStockService(id string) (*model.Stock, error) {
	return repository.GetStock(id)
}

func UpdateStockService(id string, updateReq *model.UpdateStockRequest) (*model.Stock, error) {
	stock, err := repository.GetStock(id)
	if err != nil {
		return nil, err
	}
	updateStockFields(stock, updateReq)
	updatedStock, err := repository.UpdateStock(stock)
	if err != nil {
		return nil, err
	}

	return updatedStock, nil
}

func DeleteStockService(id string) error {
	return repository.DeleteStock(id)
}

func updateStockFields(stock *model.Stock, request *model.UpdateStockRequest) {
	reqVal := reflect.ValueOf(request).Elem()
	stockVal := reflect.ValueOf(stock).Elem()
	for i := 0; i < reqVal.NumField(); i++ {
		field := reqVal.Field(i)
		if field.IsNil() {
			continue
		}
		stockField := stockVal.FieldByName(reqVal.Type().Field(i).Name)
		if stockField.IsValid() && stockField.CanSet() {
			stockField.Set(field.Elem())
		}
	}
}
