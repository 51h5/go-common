package utils

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

// 创建新的 Background Context, 并继承 ctx 中的 trace 信息
func ContextWithSpan(ctx context.Context) context.Context {
	return ContextWithSpanBy(ctx, context.Background())
}

// 给指定的 parent context 附加 ctx 中的 trace 信息
func ContextWithSpanBy(ctx, parent context.Context) context.Context {
	return trace.ContextWithSpan(parent, trace.SpanFromContext(ctx))
}

// 创建新的 支持超时 的 Background Context, 并继承 ctx 中的 trace 信息
func ContextWithSpanAndTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ContextWithSpan(ctx), timeout)
}
