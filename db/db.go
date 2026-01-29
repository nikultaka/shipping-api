package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"shipping-api/models"
	"shipping-api/models/tracking"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Printf("⚠️ Failed to connect to MySQL: %v", err)
		log.Println("⚠️ Falling back to in-memory database for now...")
		initInMemoryDB()
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("✅ Successfully connected to MySQL database!")

	autoMigrate()
}

func autoMigrate() {
	err := DB.AutoMigrate(
		&models.Order{},
		&tracking.TrackingEvent{},
	)
	if err != nil {
		log.Println("⚠️ Auto migration failed:", err)
	} else {
		fmt.Println("✅ Database tables created/updated")
	}
}

var inMemoryOrders = make(map[string]interface{})

func initInMemoryDB() {
	fmt.Println("📦 Using in-memory database (fallback mode)")
}

func OrderExists(orderID string) bool {
	if DB != nil {
		return false
	}

	_, exists := inMemoryOrders[orderID]
	return exists
}
