package config

import (
	"fmt"
	"frs-planning-backend/db"
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/api/routes"
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type RestConfig struct {
	server *gin.Engine
}

func NewRest() RestConfig {
	db := db.New()
	app := gin.Default()
	server := NewRouter(app)
	middleware := middleware.New(db)

	var (
		//=========== (PACKAGE) ===========//
		// mailerService mailer.Mailer         = mailer.New()

		//=========== (REPOSITORY) ===========//
		userRepository repository.UserRepository = repository.NewUser(db)

		//=========== (SERVICE) ===========//
		authService service.AuthService = service.NewAuth(userRepository, db)
		userService service.UserService = service.NewUser(userRepository, db)

		//=========== (CONTROLLER) ===========//
		authController controller.AuthController = controller.NewAuth(authService)
		userController controller.UserController = controller.NewUser(userService)
	)

	routes.Auth(server, authController, middleware)
	routes.User(server, userController, middleware)
	return RestConfig{
		server: server,
	}
}

func (ap *RestConfig) Start() {
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")
	if port == "" {
		port = "8998"
	}

	serve := fmt.Sprintf("%s:%s", host, port)
	if err := ap.server.Run(serve); err != nil {
		log.Panicf("failed to start server: %s", err)
	}
	log.Println("server start on port ", serve)
}
