package handler

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/pkg/ginx"
)

func (h *Handler) Ping(c *gin.Context) {
	ginx.Res(c, "Pong", nil)
}
