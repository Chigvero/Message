package handler

import (
	Intern "github.com/Chigvero/Messageio"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (h *Handler) CreateMessage(c *gin.Context) {
	var msg Intern.Message
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println(err)
		res := Error{
			"Please enter correct data type",
		}
		c.AbortWithStatusJSON(400, res)
		return
	}
	id, err := h.service.CreateMessage(msg)
	if err != nil {
		log.Println(err)
		res := Error{
			"Please try again later",
		}
		c.AbortWithStatusJSON(500, res)
		return
	}
	msg.Id = id
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) ProcessMessage(c *gin.Context) {
	//return s.repos.ProcessMessage()
}

func (h *Handler) GetMessageById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		res := Error{Message: "Please enter correct id"}
		c.AbortWithStatusJSON(400, res)
		return
	}
	msg, err := h.service.GetMessageById(id)
	if err != nil {
		log.Println(err)
		res := Error{Message: "Message with this id not found"}
		c.AbortWithStatusJSON(200, res)
		return
	}

	c.JSON(200, msg)
}
