package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrat-muslim/students-api/api/v1"
	"github.com/ibrat-muslim/students-api/config"
	"github.com/ibrat-muslim/students-api/storage"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/ibrat-muslim/students-api/api/docs" // for swagger
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title           Swagger for student api
// @version         1.0
// @description     This is a student service api.
// @host      localhost:8000
// @BasePath  /v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	apiV1.POST("/students", handlerV1.CreateStudent)
	apiV1.GET("/students", handlerV1.GetStudents)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
