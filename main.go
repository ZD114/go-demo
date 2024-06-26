package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"net/http"
	"zhangda.com/go-demo/config"
	"zhangda.com/go-demo/log"
	"zhangda.com/go-demo/router"
)

func main() {
	log.InitLogger(config.Config)
	config.InitDb(config.Config)

	newRouter := router.NewRouter(config.Config)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Config.GetInt("server.port")),
		Handler:      newRouter,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	log.Logger.Info("服务初始化成功")

	if err := s.ListenAndServe(); err != nil {
		log.Logger.Error("服务器启动异常！", zap.Error(err))
	}
}
