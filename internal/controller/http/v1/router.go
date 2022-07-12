// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ironsail/whydah-go-clean-template/internal/entity"
	"github.com/ironsail/whydah-go-clean-template/internal/usecase"
	"github.com/ironsail/whydah-go-clean-template/pkg/postgres"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs, must have in order to be able to display Swagger doc
	_ "github.com/ironsail/whydah-go-clean-template/docs"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /api/v1
func NewRouter(handler *gin.Engine, pg postgres.Postgres) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/api/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// TODO need 1 more health api to check status of dependencies (db, external services ...)

	// Routers
	h := handler.Group("/api/v1")
	{
		// Use case
		userUseCase := usecase.NewUserUseCase(entity.Users(pg))

		newUserRoutes(h, userUseCase)
	}
}
