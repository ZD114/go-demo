package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zhangda.com/go-demo/object"
	"zhangda.com/go-demo/service"
)

func Test(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	res, err := service.TestService.Test(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))

	} else {
		ctx.JSON(http.StatusOK, res)
	}
}
