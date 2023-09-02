package models

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres" // You'll need the PostgreSQL driver for gorm
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert" // Using a testing assertion library for cleaner assertions
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func TestMigrateBooks(t *testing.T) {
	// Set up a test database
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
		t.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Run migrations
	err = MigrateBooks(db)
	if err != nil {
		t.Fatalf("Error running migrations: %v", err)
	}
	
	// Perform assertions to verify migrations were successful
	assert.True(t, db.Migrator().HasTable(&Books{}), "Books table should exist")
}

// You can write more tests for the Books struct and other functions here
