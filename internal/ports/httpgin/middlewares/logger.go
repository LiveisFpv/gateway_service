package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		entry := log.WithFields(logrus.Fields{
			"msg":       "Request recived",
			"client_ip": c.ClientIP(),
			"method":    c.Request.Method,
			"header":    c.Request.Header,
			"body":      c.Request.Body,
			"path":      c.Request.RequestURI,
			"proto":     c.Request.Proto,
			"time":      time.Now(),
		})
		entry.Info()
		c.Next()
	}
}
func ResponseLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		start := time.Now()
		// Process Request
		c.Next()
		// Stop timer
		latency := time.Since(start)

		entry := log.WithFields(logrus.Fields{
			"client_ip": c.ClientIP(),
			"duration":  latency,
			"method":    c.Request.Method,
			"header":    c.Request.Header,
			"body":      c.Request.Body,
			"path":      c.Request.RequestURI,
			"proto":     c.Request.Proto,
			"status":    c.Writer.Status(),
			"msg":       "Response sent",
		})
		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
