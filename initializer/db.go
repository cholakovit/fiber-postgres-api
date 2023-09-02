package initializer

import (
	
	"fmt"
	"log"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func Connect() (*gorm.DB, error) {
	var err error

	// config := Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Password: os.Getenv("DB_PASS"),
	// 	User:     os.Getenv("DB_USER"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// 	DBName:   os.Getenv("DB_NAME"),
	// }

	config := Config{
		Host:     "localhost",
		Port:     "5432",
		Password: "srichinmoy",
		User:     "postgres",
		SSLMode:  "disable",
		DBName:   "fiber_demo",
	}

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Kolkata",
	config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	
	// Get the underlying *sql.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get *sql.DB instance:", err)
	}

	// Perform a ping to verify the connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	fmt.Println("Connected to the database successfully!")

	return db, nil
}