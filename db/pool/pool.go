package pool

import (
	"fmt"
	"log"
	"rbac/db/config"
	"rbac/pkg/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
)


var DB *gorm.DB

func rbac(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		return fmt.Errorf("could not open db : %w", err)
	}
	
	
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("could not get sql.DB: %w", err)
	}

	// Configure connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Successfully connected to the database")
	
	err = DB.AutoMigrate(models.RegisterModels...)
	if err != nil {
		return fmt.Errorf("AutoMigrate failed: %w", err)
	}

	log.Println("GORM models mapped to migrated tables")
	return nil
}

func Disrbac(){
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("Error getting sql.DB: %v", err)
			return
		}
		
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}

