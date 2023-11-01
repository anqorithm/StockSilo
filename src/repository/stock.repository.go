// repository/stock.repository.go

package repository

import (
	"github.com/qahta0/stocksilo/config"
	"github.com/qahta0/stocksilo/model"
)

func Migrate() {
	db := config.DB()

	db.AutoMigrate(&model.Stock{})
}

func CreateStock(stock *model.Stock) (*model.Stock, error) {
	db := config.DB()
	err := db.Create(&stock).Error
	return stock, err
}

func GetStock(id string) (*model.Stock, error) {
	db := config.DB()
	var stock model.Stock
	err := db.Where("id = ?", id).First(&stock).Error
	return &stock, err
}

func GetStocks() ([]model.Stock, error) {
	db := config.DB()
	var stocks []model.Stock
	err := db.Find(&stocks).Error
	return stocks, err
}

func UpdateStock(stock *model.Stock) error {
	db := config.DB()
	return db.Save(&stock).Error
}

func DeleteStock(id string) error {
	db := config.DB()
	return db.Where("id = ?", id).Delete(&model.Stock{}).Error
}
