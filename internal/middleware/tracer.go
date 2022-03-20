package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
				global.Tracer, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
				global.Tracer, c.Request.URL.Path)
		}

		defer span.Finish()

		var spanContext = span.Context()
		var jaegerContext = spanContext.(jaeger.SpanContext)

		c.Set("X-Trace-ID", jaegerContext.TraceID().String())
		c.Set("X-Span-ID", jaegerContext.SpanID().String())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
