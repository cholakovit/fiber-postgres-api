package initializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// Call the Connect function to establish a connection
	db, err := Connect()

	// Check for errors in establishing the connection
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Close the database connection after testing
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
