package stats

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

func NumGoroutine(c *gin.Context) {
	c.JSON(200, gin.H{
		"num_goroutine": runtime.NumGoroutine(),
	})
}

func MemoryStats(c *gin.Context) {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	c.JSON(200, gin.H{
		"memory_sys": m.Sys,
	})
}
