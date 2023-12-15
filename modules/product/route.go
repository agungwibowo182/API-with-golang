package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRoutes(app *fiber.App, db *gorm.DB) {
	productRepo := NewProductRepository(db)
	productService := NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	// Define product routes
	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products", productHandler.GetAllProducts)
	app.Get("/products/:id", productHandler.GetProductByID)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)
}
