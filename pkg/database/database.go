package database

import (
	"fmt"

	"calc_example/internal/config"
	"calc_example/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func New(cfg config.DatabaseConfig) (*Database, error) {
	var db *gorm.DB
	var err error

	switch cfg.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.DBName+".db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		return nil, fmt.Errorf("неподдерживаемый драйвер базы данных: %s", cfg.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Автоматическая миграция моделей
	if err := db.AutoMigrate(
		&model.Issue{},
	); err != nil {
		return nil, fmt.Errorf("ошибка миграции базы данных: %w", err)
	}

	return &Database{db}, nil
}

func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
} 