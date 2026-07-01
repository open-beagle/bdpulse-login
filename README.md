# go-login

Go 语言实现的 OAuth 登录中间件库，支持多种 Git 托管平台的身份认证。

## 支持的平台

| 平台                     | 协议     | 说明                                 |
| ------------------------ | -------- | ------------------------------------ |
| GitHub                   | OAuth2   | 支持 GitHub.com 和 GitHub Enterprise |
| GitLab                   | OAuth2   | 支持 GitLab.com 和自托管实例         |
| Gitea                    | OAuth2   | 自托管 Git 服务                      |
| Gogs                     | 表单认证 | 自托管 Git 服务                      |
| Bitbucket                | OAuth2   | Bitbucket Cloud                      |
| Bitbucket Server (Stash) | OAuth1   | 企业版 Bitbucket                     |
| Gitee                    | OAuth2   | 码云                                 |
| Beagle                   | OAuth2   | Beagle 平台                          |

## 安装

```bash
go get github.com/open-beagle/bdpulse-login
```

## 快速开始

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/open-beagle/bdpulse-login/login"
    "github.com/open-beagle/bdpulse-login/login/github"
)

func main() {
    middleware := &github.Config{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        Scope:        []string{"repo", "user"},
    }

    http.Handle("/login", middleware.Handler(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            if err := login.ErrorFrom(ctx); err != nil {
                http.Error(w, err.Error(), 500)
                return
            }
            token := login.TokenFrom(ctx)
            fmt.Fprintf(w, "Access Token: %s", token.Access)
        }),
    ))

    http.ListenAndServe(":8080", nil)
}
```

## 使用示例

运行示例程序：

```bash
go run example/main.go \
  --provider github \
  --client-id <your-client-id> \
  --client-secret <your-client-secret>
```

### 命令行参数

| 参数                     | 说明                                                                      |
| ------------------------ | ------------------------------------------------------------------------- |
| `--provider`             | 认证提供商 (github, gitlab, gitea, gogs, bitbucket, stash, gitee, beagle) |
| `--provider-url`         | 自托管服务地址 (gitea, gogs, gitlab, stash)                               |
| `--client-id`            | OAuth2 Client ID                                                          |
| `--client-secret`        | OAuth2 Client Secret                                                      |
| `--consumer-key`         | OAuth1 Consumer Key (stash)                                               |
| `--consumer-private-key` | OAuth1 RSA 私钥文件路径 (stash)                                           |
| `--redirect-url`         | OAuth 回调地址                                                            |
| `--address`              | HTTP 服务监听地址 (默认 :8080)                                            |
| `--dump`                 | 启用请求调试日志                                                          |

## 特性

- 支持 OAuth1 和 OAuth2 协议
- 可选的 OpenTelemetry 链路追踪支持
- 内置请求日志和调试工具
- 简洁的中间件 API 设计

## License

BSD-style license
