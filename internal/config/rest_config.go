package config

import (
	"fmt"
	"frs-planning-backend/db"
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/api/routes"
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/middleware"
	mailer "frs-planning-backend/internal/pkg/email"
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
		mailerService mailer.Mailer = mailer.New()

		//=========== (REPOSITORY) ===========//
		userRepository                  repository.UserRepository                  = repository.NewUserRepository(db)
		classRepository                 repository.ClassRepository                 = repository.NewClassRepository(db)
		courseRepository                repository.CourseRepository                = repository.NewCourseRepository(db)
		workspaceRepository             repository.WorkspaceRepository             = repository.NewWorkspaceRepository(db)
		workspaceCollaboratorRepository repository.WorkspaceCollaboratorRepository = repository.NewWOrkspaceCollaboratorRepository(db)

		//=========== (SERVICE) ===========//
		authService                  service.AuthService                   = service.NewAuthService(userRepository, mailerService, db)
		userService                  service.UserService                   = service.NewUserService(userRepository, db)

		classService                 service.ClassService                  = service.NewClassService(classRepository, courseRepository, db)

		classService                 service.ClassService                  = service.NewClassService(classRepository, db)

		courseService                service.CourseService                 = service.NewCourseService(courseRepository)
		workspaceService             service.WorkspaceService              = service.NewWorkspaceService(workspaceRepository, db)
		workspaceCollaboratorService service.WorskspaceCollaboratorService = service.NewWorkspaceCollaboratorService(workspaceCollaboratorRepository, userRepository, db)

		//=========== (CONTROLLER) ===========//
		authController                  controller.AuthController                  = controller.NewAuth(authService)
		userController                  controller.UserController                  = controller.NewUser(userService)
		classController                 controller.ClassController                 = controller.NewClassController(classService)
		courseController                controller.CourseController                = controller.NewCourseController(courseService)
		workspaceController             controller.WorkspaceController             = controller.NewWorkspace(workspaceService)
		workspaceCollaboratorController controller.WorkspaceCollaboratorController = controller.NewWorkspaceCOllaborator(workspaceCollaboratorService)
	)

	// Register all routes
	routes.Auth(server, authController, middleware)
	routes.User(server, userController, middleware)
	routes.Class(server, classController, middleware)

	routes.Course(server, courseController, classController, middleware)
	routes.ClassSetting(server, controller.NewClassSettingController(service.NewClassSettingService(repository.NewClassSettingRepository(db), db)), middleware)

	routes.Course(server, courseController, middleware)

	routes.Workspace(server, workspaceController, middleware)
	routes.WorkspaceCollaborator(server, workspaceCollaboratorController, middleware)
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
