package database

import (
	"Magazine_backend/pkg/utils"
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config — структура для хранения конфигурации подключения
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

// LoadConfigFromEnv загружает настройки базы данных из переменных окружения
func LoadConfigFromEnv() Config {
	cfg := Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PWD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_NAME"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	// utils.Logger.Printf("Проверка загрузки\nuser=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
	// 	cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	return cfg
}

// DB — глобальная переменная для хранения подключения к базе данных
var DbPostgres *gorm.DB
var sqlDB *sql.DB

// Connect устанавливает соединение с базой данных
func Connect(cfg Config) error {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)
	utils.Logger.Printf("Проверка подключения\n%s", dsn)
	var err error
	DbPostgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("Ошибка подключения к базе данных: %w", err)
	}

	// Проверяем соединение
	sqlDB, err := DbPostgres.DB()
	if err != nil {
		log.Fatal("Ошибка получения sql.DB: ", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("БД недоступна: ", err)
	}

	utils.Logger.Info("Успешное подключение к базе данных")
	return nil
}

func CreateObjDB(dst ...interface{}) {
	if err := DbPostgres.AutoMigrate(dst...); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
}

// Close закрывает соединение с базой данных
func Close() error {
	if DbPostgres != nil && sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			return err
		}
	}
	return nil
}
