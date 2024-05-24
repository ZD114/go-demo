package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"zhangda.com/go-demo/controller"
	"zhangda.com/go-demo/log"
	"zhangda.com/go-demo/object"
)

func NewRouter(viper *viper.Viper) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()

	if viper.GetString("spring.profiles.active") == "dev" {
		server.Use(Cors())
	}

	server.Use(ginzap.Ginzap(log.Logger, viper.GetString("logging.date-time-format"), false))
	server.Use(Recovery)

	group := server.Group("")
	{
		group.GET("/test", controller.Test)
	}

	return server
}

func Recovery(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {

			ctx.JSON(http.StatusOK, object.FailMsg("router监控到错误！"))
		}
	}()
	ctx.Next()
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		c.Next()
	}
}
