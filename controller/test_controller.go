package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhangda.com/go-demo/object"
	"zhangda.com/go-demo/service"
)

func Test(ctx *gin.Context) {
	if res, err := service.TestService.Test(); err != nil {
		ctx.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}
