package restapi

import (
	"Axsprav/internal/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


type RESTAPI struct {
	server *gin.Engine
	logger *zap.SugaredLogger
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func New(logger *zap.SugaredLogger) *RESTAPI {

	controller := controller.NewController()
	r := gin.New()

	r.Use(CORSMiddleware())

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	//r.LoadHTMLGlob("./templates/*")

	v1 := r.Group("/api/v1")
	{
		v1.GET("UpdateInventJournalTable", controller.UpdateInventJournalTable)

	}

	return &RESTAPI{server: r, logger: logger}
}

func (rapi *RESTAPI) Start(port int) {

	rapi.server.Run(fmt.Sprintf(":%v", port))

}

