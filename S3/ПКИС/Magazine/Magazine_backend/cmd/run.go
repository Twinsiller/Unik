package cmd

import (
	v1 "Magazine_backend/internal/api/v1"
	"Magazine_backend/internal/database"
	"Magazine_backend/internal/models"
	"Magazine_backend/pkg/utils"

	"github.com/joho/godotenv"
)

func Run() error {
	// Запуск логгера
	utils.InitLogger("pkg/utils/app.log")

	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		utils.Logger.Fatalf("Error loading .env file")
		return err
	}
	if err := database.Connect(database.LoadConfigFromEnv()); err != nil {
		return err
	}
	defer database.Close()

	utils.Logger.Info("Передача моделей")
	database.CreateObjDB(
		&models.User{},
		&models.Supplier{},
		&models.StockMovement{},
		&models.Stock{},
		&models.Purchase{},
		&models.PurchaseItem{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Manufacturer{},
		&models.Category{},
	)
	utils.Logger.Info("Успешная передача моделей")

	v1.Apies()

	return nil
}
