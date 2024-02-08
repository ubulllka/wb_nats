package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ubulllka/wbL0/internal/config"
	"github.com/ubulllka/wbL0/internal/logger"
	"github.com/ubulllka/wbL0/internal/models"
)

var DB *gorm.DB

func InitializeDB() error {
	var err error
	logg := logger.GetLogger()
	info := config.CONFIG.DB
	urlPostgres := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.Name)

	DB, err = gorm.Open("postgres", urlPostgres)
	if err != nil {
		log.Panic(err)
		logg.Panic(err)
		return err
	}

	DB.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Item{}, &models.Payment{})

	err = DB.DB().Ping()
	if err != nil {
		logg.Panic(err)
		return err
	}
	logg.Info("Init database")

	return nil
}

func AddDbOrder(order models.Order) error {
	logg := logger.GetLogger()
	tx := DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		logg.Panic(err)
		tx.Rollback()
	}

	tx.Commit()
	return nil
}
