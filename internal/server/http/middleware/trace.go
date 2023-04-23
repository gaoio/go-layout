package middleware

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/pkg/logx"
	"go-layout/internal/pkg/utils"
)

func (m *Middleware) Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		trace := utils.UUID()
		c.Set("trace", trace)
		ctx = logx.NewContext(ctx, logx.String("trace", trace))
		ctx = logx.NewContext(ctx, logx.String("ip", c.ClientIP()))
		c.Request = c.Request.WithContext(ctx)
		logx.WithContext(ctx).Info("access",
			logx.String("method", c.Request.Method),
			logx.String("path", c.Request.URL.Path),
		)

		c.Next()
	}
}
