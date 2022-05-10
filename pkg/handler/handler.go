package handler

import (
	"twittie/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services  *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth") 
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api") 
	{
		posts := api.Group("/posts") 
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getAllPosts)
			posts.GET("/:id", h.getPostById)
			posts.DELETE("/:id", h.deletePost)
		}
	}
	return router
}