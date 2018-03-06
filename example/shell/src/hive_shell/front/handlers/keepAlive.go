package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hive_shell/config"
	"hive_shell/worker"
	"vega/go-kit/actor/registry"
)

func KeepAlive(c *gin.Context) {

	actor, err := registry.Whereis(config.SupervisorName, config.WorkerActorName)
	if err != nil {
		c.JSON(503, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	err = actor.(*worker.Worker).KeepAlive()
	if err != nil {
		c.JSON(503, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
	return
}
