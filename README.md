# gf2-authz - GoFrame v2 + GORM + React + Material UI

**Read this in other languages: [English Docs](README.md) | [Chinese / 中文文档](README.zh-CN.md)

gf2-auth is a fork of [eko/authz](https://github.com/eko/authz), Backend with GoFrame instead of Fiber, frontend migrated from react-scripts to Vite.

This project brings a backend server with its frontend for managing authorizations.

You can use both Role-Based Acccess Control (RBAC) and Attribute-Based Access Control (ABAC).

## Why use it?

🌍  A centralized backend for all your applications authorizations

🙋‍♂️  Supports Role-Based Access Control (RBAC)

📌  Supports Attribute-Based Access Control (ABAC)

⚙️   Go SDKs available

✅  Reliable: Authz uses Authz itself for managing its own internal authorizations

🔍  Audit: We log each check decisions and which policy matched

🔐  Single Sign-On: Use your enterprise SSO to log into the web UI, using OpenID Connect

🕵️‍♂️  Observability: Retrieve metrics and tracing data into your prefered tools

## SDKs

In order to help you, we have the following available SDKs:

- [Go](https://github.com/vulcangz/gf2-authz/tree/main/pkg/sdk)
- [blog example](https://github.com/vulcangz/gf2-authz/tree/main/examples/blog)

Please check their documentations for detailled usage. They all use `gRPC` for communicating with the Authz backend (server-to-server).

## Getting started

To get started with this project, run

### Running with default config

No configuration is required. 

#### STEP 1: Backend

```bash
  git clone https://github.com/vulcangz/gf2-authz.git
  cd gf2-authz
  go mod tidy
  go run main.go
```

The system then runs with the default config, using SQLite in-memory database.



#### STEP 2: Admin Dashboard(UI)

```bash
  cd ui
  pnpm i
  pnpm dev
```

visiting http://localhost:3000

Sign in with default credentials: `admin` / `changeme`.

#### STEP 3: Examples blog

1. Create a service account under menu `Service accounts` in dashboard.
2. Edit [main.go](https://github.com/vulcangz/gf2-authz/blob/main/examples/blog/main.go#L18-L19). Replace the `client_id`, `client_secret` which obtained in the previous step.
3. Edit the principal(name: `auth-sa-(your service account name)`). Assign a role(`authz-admin`) to it.
4. Run the test:
```
go run main.go
```
5. visiting [metrics api](http://localhost:8080/v1/metrics) for Prometheus metrics observability(default config: disable). 

### Running with your config

Save [example config](https://github.com/vulcangz/gf2-authz/blob/main/manifest/config/config.example.yaml) as `config.yaml`. Edit it with your config.

Then, same steps as above.


that's all you need to get started!


## Testing

- When testing with a MySQL database, "29 scenarios (29 passed). 256 steps (256 passed)";
- When testing with a PostgreSQL database, "29 scenarios (29 passed). 256 steps (256 passed)";
- When testing with an SQLite database, "29 scenarios (29 passed). 256 steps (256 passed)".

On a development machine (Intel® Core™ i5-4570 CPU, virtual machine, 2 cores, 10 GB RAM), a full features test took approximately 23–25 seconds for MySQL, 13–14 seconds for PostgreSQL, and 7–8 seconds for SQLite.

The following tests use a MySQL database.

1. Save [example config](https://github.com/vulcangz/gf2-authz/blob/main/manifest/config/config.example.yaml) as `config.yaml`. Edit it with your config.
2. Create DB you defined in the config.yaml.
3. Run the test:
```
$ export GF_GMODE=testing

$ go test -count=1 --tags=functional -v ./functional

# or just test feature "action", other features: check, compiled, policy, principal, resource, role, user
$ go test -count=1 --tags=functional -v ./functional -t @action
```
The test results are similar to the following:
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

## Acknowledgments

- [eko/authz](https://github.com/eko/authz)
- [GoFrame](https://github.com/gogf/gf)
- [Gothic](https://github.com/jrapoport/gothic): methods TruncateAll(), DropAll()
