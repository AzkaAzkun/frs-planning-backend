package config

import (
	"film-management-api-golang/db"
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/api/routes"
	"film-management-api-golang/internal/api/service"
	"film-management-api-golang/internal/middleware"
	"fmt"
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
		userRepository     repository.UserRepository     = repository.NewUser(db)
		genreRepository    repository.GenreRepository    = repository.NewGenre(db)
		filmRepository     repository.FilmRepository     = repository.NewFilm(db)
		filmListRepository repository.FilmListRepository = repository.NewFilmList(db)
		reviewRepository   repository.ReviewRepository   = repository.NewReview(db)
		reactionRepository repository.ReactionRepository = repository.NewReaction(db)

		//=========== (SERVICE) ===========//
		authService     service.AuthService     = service.NewAuth(userRepository, db)
		userService     service.UserService     = service.NewUser(userRepository, db)
		genreService    service.GenreService    = service.NewGenre(genreRepository, db)
		filmService     service.FilmService     = service.NewFilm(filmRepository, genreRepository, db)
		filmListService service.FilmListService = service.NewFilmList(filmListRepository, filmRepository, reviewRepository, db)
		reviewService   service.ReviewService   = service.NewReview(reviewRepository, filmRepository, db)
		reactionService service.ReactionService = service.NewReaction(reactionRepository, reviewRepository, db)

		//=========== (CONTROLLER) ===========//
		authController     controller.AuthController     = controller.NewAuth(authService)
		userController     controller.UserController     = controller.NewUser(userService)
		genreController    controller.GenreController    = controller.NewGenre(genreService)
		filmController     controller.FilmController     = controller.NewFilm(filmService)
		filmListController controller.FilmListController = controller.NewFilmList(filmListService)
		reviewController   controller.ReviewController   = controller.NewReview(reviewService)
		reactionController controller.ReactionController = controller.NewReaction(reactionService)
	)

	routes.Auth(server, authController, middleware)
	routes.User(server, userController, middleware)
	routes.Genre(server, genreController, middleware)
	routes.Film(server, filmController, middleware)
	routes.FilmList(server, filmListController, middleware)
	routes.Review(server, reviewController, middleware)
	routes.Reaction(server, reactionController, middleware)
	return RestConfig{
		server: server,
	}
}

func (ap *RestConfig) Start() {
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")
	if port == "" {
		port = "8090"
	}

	serve := fmt.Sprintf("%s:%s", host, port)
	if err := ap.server.Run(serve); err != nil {
		log.Panicf("failed to start server: %s", err)
	}
	log.Println("server start on port ", serve)
}
