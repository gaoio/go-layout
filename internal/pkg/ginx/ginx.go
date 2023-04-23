package ginx

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/pkg/errors"
	"net/http"
)

func resJson(c *gin.Context, code int, message string, data interface{}) {
	var trace string
	val, exist := c.Get("trace")
	if exist {
		trace = val.(string)
	}

	obj := gin.H{
		"code":  code,
		"msg":   message,
		"data":  data,
		"trace": trace,
	}

	c.JSON(http.StatusOK, obj)
}

func Res(c *gin.Context, data interface{}, err error) {
	if err != nil {
		e, ok := err.(*errors.ResponseError)
		if !ok {
			resJson(c, 500, "服务器出了点小错", "")
		} else {
			resJson(c, e.Code, e.Message, "")
		}
		return
	}
	resJson(c, 0, "ok", data)
}
