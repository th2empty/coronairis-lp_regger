package handler

import (
	"coronairis-lp_web-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/images", "web/images")
	router.Static("/styles", "web/styles")
	router.Static("/scripts", "web/scripts")

	router.GET("/products/bots/cbot/v1/registration", h.GetIndexPage)

	api := router.Group("/api/v1")
	{
		api.POST("/register", h.Register)
		api.POST("/update", h.UpdateToken)
	}

	return router
}
