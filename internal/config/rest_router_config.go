package config

import (
	"fmt"
<<<<<<< HEAD
=======
	"frs-planning-backend/internal/api/routes"
>>>>>>> b258106 (Commit fitur CRUD Course and Classes)
	"frs-planning-backend/internal/middleware"
	mylog "frs-planning-backend/internal/pkg/logger"
	"frs-planning-backend/internal/pkg/response"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
<<<<<<< HEAD
)

func NewRouter(server *gin.Engine) *gin.Engine {
=======
	"gorm.io/gorm"
)

func NewRouter(server *gin.Engine, db *gorm.DB) *gin.Engine {
>>>>>>> b258106 (Commit fitur CRUD Course and Classes)
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Route Not Found",
		})
	})

	server.MaxMultipartMemory = 30 * 1024 * 1024
	server.Use(customRecovery())
	server.Use(middleware.CORSMiddleware())

	server.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

<<<<<<< HEAD
=======
	// Register routes for classes and courses
	routes.RegisterClassRoutes(server, db)
	routes.RegisterCourseRoutes(server, db)

>>>>>>> b258106 (Commit fitur CRUD Course and Classes)
	return server
}

func customRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var wrappedErr error
				if e, ok := err.(error); ok {
					wrappedErr = e
				} else {
					wrappedErr = fmt.Errorf("%v", err)
				}

				fmt.Println(mylog.ColorizePanic(fmt.Sprintf("\n[Recovery] Panic occurred: %v\n", err)))
				stack := debug.Stack()
				coloredStack := mylog.ColorizePanic(string(stack))

				fmt.Fprintln(os.Stderr, coloredStack)
				response.NewFailed("server panic occured", wrappedErr).
					SendWithAbort(ctx)
			}
		}()

		ctx.Next()
	}
}
