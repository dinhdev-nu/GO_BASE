package routers

import (
	c "github.com/dinhdev-nu/GO_BASE/internal/controller"
	"github.com/gin-gonic/gin"
)


func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/api")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong)

	}
	return r
}

