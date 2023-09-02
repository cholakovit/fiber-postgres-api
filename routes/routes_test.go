package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRoutes(t *testing.T) {
	t.Run("Routes Testing", func(t *testing.T) {
		app := fiber.New()

		// Call the route f-n to set up routes
		Routes(app) 

		// Create a test request and send it to the app
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Expected nil error, but got %v", err)
		}

		// Check the response status code
		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("Expected status code %d, but got %d", fiber.StatusOK, resp.StatusCode)
		}

		// TODO: Check the response body or other expectations as needed
		// For now, let's assume you're checking for a JSON response with a message field

		// expectedJSON := `{"message":"books fetched successfully"}`
		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	t.Fatalf("Failed to read response body: %v", err)
		// }
		// if string(body) != expectedJSON {
		// 	t.Fatalf("Expected response body %s, but got %s", expectedJSON, string(body))
		// }
	})
}