// config/router.go

package config

import (
	"os"
	"tugas-bootcamp/modules/cart"
	"tugas-bootcamp/modules/product"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	product.ProductRoutes(app, db)
	cart.CartRoutes(app, db)

	// Add other module routes as needed
	// ...

	// Dynamic configuration for port and app listening
	port := os.Getenv("PORT")
	portStr := ":" + port
	app.Listen(portStr)
}
