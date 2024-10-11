package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware() fiber.Handler {
	origins := os.Getenv("ALLOWED_ORIGIN")

	if origins == "" {
		origins = "*"
		log.Println("Allowed Origin is not set, defaulting to '*'")
	}

	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		AllowMethods:     "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE",
	})
}
