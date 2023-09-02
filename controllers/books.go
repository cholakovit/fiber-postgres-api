package controllers

import (
	"context"
	"fiber-postgres-api/models"
	"fiber-postgres-api/queries"
	"fmt"
	"time"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var 	DB *gorm.DB
// Improving the performance of fetching books using concurrent execution.
func GetBooks(c *fiber.Ctx) error {
	// Create a context for managing the goroutine
	ctx, cancel := context.WithCancel(c.Context())
	defer cancel()

	// Create a channel to receive the result or error from the goroutine
	booksCh := make(chan []models.Books)
	errCh := make(chan error)
	defer close(booksCh)
	defer close(errCh)

	// Start a goroutine to fetch books
	go func() {
		books, err := queries.GetBooksQuery()
		if err != nil {
			errCh <- err
			return
		}
		booksCh <- books
	}()

	select {
		case books := <- booksCh:
			c.Status(http.StatusOK).JSON(&fiber.Map{
				"message": "books fetched successfully",
				"data":    books,
			})
		case err := <-errCh:
			c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": err.Error()})
		case <- ctx.Done():
			// Handle cancellation
			c.Status(http.StatusRequestTimeout).JSON(
				&fiber.Map{"message": "request canceled or timed out"},
			)
	}
	return nil
}

// While this approach might not provide a huge performance improvement for fetching a single book, it demonstrates how to handle 
// concurrency and cancellation within the context of my Fiber application.
func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Id cannot be empty",
		})
		return nil
	}

	// Creating a context for managing the goroutine
	ctx, cancel := context.WithCancel(c.Context())
	defer cancel()

	// Create a channel to receive the book or error from the goroutine
	bookCh := make(chan models.Books)
	errCh := make(chan error)
	defer close(bookCh)
	defer close(errCh)

	// Start a goroutine to fetch the book by ID
	go func() {
		book, err := queries.GetBookByIDQuery(id)
		if err != nil {
			errCh <- err
			return
		}
		select {
			case <-ctx.Done():
				return // return early if context is canceled
			case bookCh <- *book:
		}
	}()

	select {
		case book := <-bookCh:
			c.Status(http.StatusOK).JSON(&fiber.Map{
				"message": "book ID fetched successfully",
				"data":    book,
			})
		case err := <- errCh:
			c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": err.Error()})
		case <- ctx.Done():
			// Handle cancelation
			c.Status(http.StatusRequestTimeout).JSON(
				&fiber.Map{"message": "request canceled or timed out"},
			)
	}

	return nil
}

// Using goroutines in the CreateBook function might not be the best approach in this scenario. Goroutines are generally more useful 
// when you have multiple tasks that can be executed concurrently to improve performance.

func CreateBook(c *fiber.Ctx) error {
	book := &models.Books{}
	err := c.BodyParser(book)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	// Create a buffered channel for sending notifications
	notificationCh := make(chan string, 1)
	defer close(notificationCh)

	// Start a goroutine to send notifications
	go func() {
		sendNotification(book.Title, notificationCh)
	}()

	err = queries.CreateBookQuery(book)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "failed to create book"})
		return err
	}

	// Wait for notification result
	select {
	case notification := <- notificationCh:
		c.Status(http.StatusOK).JSON(&fiber.Map{
			"message":      "book has been added",
			"notification": notification,
		})
	}

	return nil
}

func sendNotification(title string, ch chan<- string) {
	// Simulate sending a notification
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("Notification: Book '%s' has been created", title)
}

// In the case of the DeleteBook function, using goroutines might not provide significant advantages. Deletion is often a 
// straightforward operation, and using goroutines might overcomplicate the code without providing a clear benefit in 
// terms of performance or parallelism.

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	// Create a buffered channel for logging
	logCh := make(chan string, 1)
	defer close(logCh)

	// Start a goroutine to log the deletion
	go func() {
		logDeletion(id, logCh)
	}()

	err := queries.DeleteBookQuery(id)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete book"})
		return err
	}

	// Wait for the log result
	select {
	case logMessage := <-logCh:
		fmt.Println(logMessage)
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book deleted successfully",
	})

	return nil
}

func logDeletion(id string, ch chan<- string) {
	// Simulate logging the deletion
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("Deleted book with ID: %s", id)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	// Parse the request body into a map
	var requestBody map[string]interface{}
	if err := c.BodyParser(&requestBody); err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid request body",
		})
		return err
	}

	// Create a buffered channel for logging
	logCh := make(chan string, 1)
	defer close(logCh)

	// Start a goroutine to log the update
	go func() {
		logUpdate(id, logCh)
	}()

	err := queries.UpdateBookQuery(requestBody, id)
	if err != nil {
		c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": err.Error(),
		})
		return err
	}

	// Wait for the log result
	select {
	case logMessage := <-logCh:
		fmt.Println(logMessage)
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book updated successfully",
	})
	return nil
}

func logUpdate(id string, ch chan<- string) {
	// Simulate logging the update
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("Updated book with ID: %s", id)
}




// func UpdateBook(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	if id == "" {
// 		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "id cannot be empty",
// 		})
// 		return nil
// 	}

// 	// Parse the request body into a map
// 	var requestBody map[string]interface{}
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "invalid request body",
// 		})
// 		return err
// 	}

// 	err := queries.UpdateBookQuery(requestBody, id)
// 	if err != nil {
// 		c.Status(http.StatusOK).JSON(&fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "book updated successfully",
// 	})
// 	return nil
// }