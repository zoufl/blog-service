package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (a AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := a.body.Write(p); err != nil {
		return n, err
	}

	return a.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyWriter

		beginTime := time.Now().Unix()
		ctx.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		global.Logger.WithFields(fields).Infof(ctx, "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			ctx.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
