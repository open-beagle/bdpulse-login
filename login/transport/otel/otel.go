package otel

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Transport 返回一个包装了 OpenTelemetry 追踪的 http.RoundTripper
// 用于追踪所有 OAuth 登录相关的 HTTP 请求
func Transport(base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return otelhttp.NewTransport(base)
}
