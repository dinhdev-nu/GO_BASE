package controller

import (
	"net/http"

	"github.com/dinhdev-nu/GO_BASE/internal/service"
	"github.com/gin-gonic/gin"
)

type PongController struct {
	ps *service.PongService
} // Create a struct to hold the controller

func NewPongController() *PongController{
	return &PongController{
		ps: service.NewPongService(),
	}
} // Create a function to instantiate the controller


// Pong is a function to return a pong message
func (pc *PongController) Pong(c *gin.Context){
	name:= c.Param("name")
	id:= c.DefaultQuery("id", "1")
	c.JSON((http.StatusOK), gin.H{
		"message": pc.ps.GetInforPong(),
		"name": name,
		"id": id,
	})
}

