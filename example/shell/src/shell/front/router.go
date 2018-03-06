package front

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shell/front/handlers"
	"shell/front/handlers/stats"
)

func NewRouter(config *HttpdParams) *gin.Engine {
	router := gin.Default()

	router.GET("/v1/ping", handlers.Ping)

	router.GET(routerStr(config.ShellName, "status"), handlers.Status)
	router.GET(routerStr(config.ShellName, "keep_alive"), handlers.KeepAlive)

	router.POST(routerStr(config.ShellName, "stop"), handlers.Stop)

	// stats
	router.GET("/v1/runtime/goroutine", stats.NumGoroutine)
	router.GET("/v1/runtime/memory", stats.MemoryStats)

	return router
}

func routerStr(shellName, action string) string {
	return fmt.Sprintf("/v1/%s/%s", shellName, action)
}
