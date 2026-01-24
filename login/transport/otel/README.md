# OpenTelemetry Transport for go-login

这个包为 `go-login` OAuth 登录库提供 OpenTelemetry 追踪支持。

## 使用方法

### GitLab OAuth 登录

```go
import (
    "net/http"

    "github.com/drone/go-login/login/gitlab"
    "github.com/drone/go-login/login/transport/otel"
)

// 创建 GitLab OAuth 中间件
middleware := gitlab.New(
    gitlab.WithClient(&http.Client{
        Transport: otel.Transport(http.DefaultTransport),
    }),
)
```

### GitHub OAuth 登录

```go
import (
    "net/http"

    "github.com/drone/go-login/login/github"
    "github.com/drone/go-login/login/transport/otel"
)

middleware := github.New(
    github.WithClient(&http.Client{
        Transport: otel.Transport(http.DefaultTransport),
    }),
)
```

### Gitea OAuth 登录

```go
import (
    "net/http"

    "github.com/drone/go-login/login/gitea"
    "github.com/drone/go-login/login/transport/otel"
)

middleware := gitea.New(
    gitea.WithClient(&http.Client{
        Transport: otel.Transport(http.DefaultTransport),
    }),
)
```

## 追踪效果

启用后，所有 OAuth 相关的 HTTP 请求都会在 Jaeger 中显示：

```
▼ GET /login                               [500ms]
  │
  ├─▶ HTTP GET                             [200ms]
  │   http.url: https://gitlab.com/oauth/authorize
  │   http.status_code: 302
  │
  └─▶ HTTP POST                            [150ms]
      http.url: https://gitlab.com/oauth/token
      http.status_code: 200
```

## 依赖

```bash
go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
```

## 更新 go.mod

在 `.tmp/go-login/go.mod` 中添加：

```go
require go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.64.0
```

## 注意事项

- 确保在应用启动时初始化了 OpenTelemetry TracerProvider
- 所有 OAuth 提供商（GitLab、GitHub、Gitea、Bitbucket 等）都支持
- 追踪包括 OAuth 授权请求和 Token 交换请求
