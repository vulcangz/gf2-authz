# gf2-authz - GoFrame v2 + GORM + React + Material UI

**其他语言版本: [English Docs](README.md) | [Chinese / 中文文档](README.zh-CN.md)

gf2-auth 是 [eko/authz](https://github.com/eko/authz) 的复刻版。后端使用 GoFrame 框架替代了 Fiber，前端构建工具从 react-scripts 迁移到 Vite。

该项目提供了一个包含前端的前端后端服务器，用于管理授权。

您可以同时使用基于角色的访问控制（RBAC）和基于属性的访问控制（ABAC）。

## 为什么要使用它？

🌍  为所有应用程序的授权提供集中式后端

🙋‍♂️  支持基于角色的访问控制（RBAC）

📌  支持基于属性的访问控制（ABAC）

⚙️  提供 Go SDK

✅  可靠：Authz 自身采用 Authz 来管理其内部授权

🔍  审计：我们会记录每次验证决策及匹配的策略

🔐  单点登录：通过 OpenID Connect 使用您的企业 SSO 登录 Web 界面

🕵️‍♂️  可观测性：将指标和追踪数据导入您偏好的工具

## SDKs

为了向您提供帮助，我们提供了以下 SDK：

- [Go](https://github.com/vulcangz/gf2-authz/tree/main/pkg/sdk)
- [blog example](https://github.com/vulcangz/gf2-authz/tree/main/examples/blog)

请查阅相关文档以了解详细用法。它们均使用 `gRPC` 与 Authz 后端进行通信（服务器间通信）。

## 入门指南

要开始使用本项目，请运行

### 使用默认配置运行

无需任何配置。

#### STEP 1: Backend

```bash
  git clone https://github.com/vulcangz/gf2-authz.git
  cd gf2-authz
  go mod tidy
  go run main.go
```

随后，系统将使用默认配置运行，并采用 SQLite 内存数据库。



#### STEP 2: 管理员控制台（UI）

```bash
  cd ui
  pnpm i
  pnpm dev
```

访问 http://localhost:3000

使用默认凭据登录： `admin` / `changeme`.

#### STEP 3: 博客示例

在仪表盘的“服务账户”菜单下创建一个服务账户。
编辑 main.go 文件。将 client_id 和 client_secret 替换为上一步中获取的值。
编辑 principal（名称：auth-sa-(您的服务账户名称)）。为其分配一个角色（authz-admin）。
运行测试：

1. 在控制台 `Service accounts` 菜单下创建一个服务账户。
2. 编辑 [main.go](https://github.com/vulcangz/gf2-authz/blob/main/examples/blog/main.go#L18-L19) 文件。将 `client_id` 和 `client_secret` 替换为上一步中获取的值。
3. 编辑 principal(name: `auth-sa-(您的服务账户名称)`)。为其分配一个角色 (`authz-admin`)。
4. 运行测试:
```
go run main.go
```
5. 访问 Prometheus 指标可观测性的 [metrics api](http://localhost:8080/v1/metrics) (默认配置：disable)。

### 使用您的配置运行

将 [示例配置](https://github.com/vulcangz/gf2-authz/blob/main/manifest/config/config.example.yaml) 保存为 `config.yaml`。根据您的配置进行编辑。

然后，按照上述步骤操作。

这就是您开始所需的全部内容！


## 测试

- 使用 MySQL 数据库测试时，结论："29 scenarios (29 passed). 256 steps (256 passed)"；
- 使用 PostgreSQL 数据库测试时，结论："29 scenarios (29 passed). 256 steps (256 passed)"；
- 使用 SQLite 数据库测试时，结论："29 scenarios (29 passed). 256 steps (256 passed)"。

在一台开发机上（Intel(R) Core(TM) i5-4570 CPU/虚拟机/2核/10GB内存），完成一次全部功能（features）测试, MySQL 用时约 23~25s，PostgreSQL 约 13~14s，SQLite 约 7~8s。

以下测试使用 MySQL 数据库。

1. 将 [示例配置](https://github.com/vulcangz/gf2-authz/blob/main/manifest/config/config.example.yaml) 保存为 `config.yaml`。根据您的配置进行编辑。
2. 创建您在 config.yaml 中定义的数据库。
3. 运行测试：
```
$ export GF_GMODE=testing

$ go test -count=1 --tags=functional -v ./functional

# or just test feature "action", other features: check, compiled, policy, principal, resource, role, user
$ go test -count=1 --tags=functional -v ./functional -t @action
```

测试结果大致如下：
```
2026-04-29T14:07:49.479Z [INFO] database.go:74: mysql database is alive!

2026-04-29T14:07:49.493Z [INFO] checkAlreadyInitialized update password ok.
2026-04-29T14:07:49.494Z [INFO] {dba3bee110d9aa1830e29b2dd23aa57e} Compiler: subscribed to event dispatchers
2026-04-29T14:07:49.512Z [INFO] pid[52582]: http server started listening on [:8080]
2026-04-29T14:07:49.512Z [INFO] swagger ui is serving at address: http://127.0.0.1:8080/swagger/
2026-04-29T14:07:49.512Z [INFO] openapi specification is serving at address: http://127.0.0.1:8080/api.json
Feature: action
  Test action-related APIs
2026-04-29T14:07:50.065Z [INFO] initialize end.
...
29 scenarios (29 passed)
256 steps (256 passed)
22.787520128s
ok      github.com/vulcangz/gf2-authz/functional        23.127s
```

## 致谢

- [eko/authz](https://github.com/eko/authz)
- [GoFrame](https://github.com/gogf/gf)
- [Gothic](https://github.com/jrapoport/gothic): methods TruncateAll(), DropAll()
