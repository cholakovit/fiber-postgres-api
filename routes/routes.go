package routes

import (
	"fiber-postgres-api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/books", controllers.GetBooks)
	app.Post("/book", controllers.CreateBook)
	app.Get("/book/:id", controllers.GetBookByID)
	app.Delete("/book/:id", controllers.DeleteBook)
	app.Patch("/book/:id", controllers.UpdateBook)

}