package handler

import (
	"github.com/gin-gonic/gin"
	"start/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", h.pingHandler)

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUP)
		auth.POST("/sign-in", h.signIN)
	}

	user := api.Group("/user", h.userIdentity)
	{
		user.POST("/", h.CreateUser)
		user.GET("/:id", h.GetUserByID)
		user.GET("/check/:id", h.CheckUserByID)
		user.DELETE("/:id", h.DeleteUserByID)
		user.GET("/get_all", h.GetAllUsersIDs)
	}
	return router
}
