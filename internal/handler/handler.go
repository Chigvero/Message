package handler

import (
	"github.com/Chigvero/Messageio/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("api/v1/message")
	{
		api.POST("/", h.CreateMessage)
		api.GET("/:id", h.GetMessageById)
	}
	return router
}
