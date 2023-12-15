package cart

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CartRoutes(app *fiber.App, db *gorm.DB) {
	cartRepo := NewCartRepository(db)
	cartService := NewCartService(cartRepo)
	cartHandler := NewCartHandler(cartService)

	// Define cart routes
	app.Post("/carts", cartHandler.CreateCart)
	app.Get("/carts", cartHandler.GetAllCarts)
	app.Get("/carts/:id", cartHandler.GetCartByID)
	app.Put("/carts/:id", cartHandler.UpdateCart)
	app.Delete("/carts/:id", cartHandler.DeleteCart)
}
