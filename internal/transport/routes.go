package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/ubulllka/wbL0/internal/config"
	"github.com/ubulllka/wbL0/internal/logger"
	"github.com/ubulllka/wbL0/internal/transport/handler"
)

func InitRouter() error {
	logg := logger.GetLogger()
	confServ := config.CONFIG.Server
	router := gin.Default()
	CreateRouter(router)

	if err := router.Run(confServ.URL); err != nil {
		logg.Panic(err)
		return err
	}
	logg.Info("Router init")
	return nil
}

func CreateRouter(r *gin.Engine) {
	routes := r.Group("/orders")
	routes.GET("/", handler.GetAll)
	routes.POST("/", handler.AddFromPub)
	routes.GET("/:id", handler.Get)
}
