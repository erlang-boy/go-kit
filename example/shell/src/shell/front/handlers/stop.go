package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shell/config"
	"shell/worker"
	"go-kit/actor/registry"
)

func Stop(c *gin.Context) {
	actor, err := registry.Whereis(config.SupervisorName, config.WorkerActorName)
	if err != nil {
		c.JSON(503, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	err = actor.(*worker.Worker).Stop()
	if err != nil {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Ok",
	})
	return
}
