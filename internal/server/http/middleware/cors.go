package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) CORS() gin.HandlerFunc {
	return cors.Default()
}
