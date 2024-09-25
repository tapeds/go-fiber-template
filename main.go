package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/tapeds/go-fiber-template/cmd"
	"github.com/tapeds/go-fiber-template/config"
	"github.com/tapeds/go-fiber-template/controller"
	"github.com/tapeds/go-fiber-template/middleware"
	"github.com/tapeds/go-fiber-template/repository"
	"github.com/tapeds/go-fiber-template/routes"
	"github.com/tapeds/go-fiber-template/service"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)

		// Controller
		userController controller.UserController = controller.NewUserController(userService)
	)

	server := fiber.New()
	server.Use(middleware.CORSMiddleware())
	apiGroup := server.Group("/api")

	// routes
	routes.User(apiGroup, userController, jwtService)

	server.Static("/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Listen(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
