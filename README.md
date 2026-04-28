# gf2-authz - GoFrame v2 + GORM + React + Material UI

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

When all features are tested together, the pass rate is close to 91%. However, individual features should pass their tests. 

The following tests use a MySQL database.

1. Save [example config](https://github.com/vulcangz/gf2-authz/blob/main/manifest/config/config.example.yaml) as `config.yaml`. Edit it with your config.
2. Create DB you defined in the config.yaml.
3. Run the test:
```
$ export GF_GMODE=testing

$ go test -count=1 --tags=functional -v ./functional
or just test feature "action", other features: check, compiled, policy, principal, resource, role, user
$ go test -count=1 --tags=functional -v ./functional -t @action
```
The test results are similar to the following:
```
2026-04-25T13:09:04.026Z [INFO] database.go:74: mysql database is alive!

2026-04-25T13:09:04.032Z [INFO] initialize start...
2026-04-25T13:09:04.035Z [INFO] initializeResources start……
2026-04-25T13:09:04.108Z [INFO] initializeResources ok.
2026-04-25T13:09:04.219Z [INFO] initializePolicies ok.
2026-04-25T13:09:04.247Z [INFO] initializeRoles ok.
2026-04-25T13:09:04.267Z [INFO] initializeUser start update……
2026-04-25T13:09:04.276Z [INFO] initialize end.
2026-04-25T13:09:04.276Z [INFO] {c90e1bc7899ba918939199752b024807} Compiler: subscribed to event dispatchers
2026-04-25T13:09:04.302Z [INFO] pid[15662]: http server started listening on [:8080]
2026-04-25T13:09:04.302Z [INFO] swagger ui is serving at address: http://127.0.0.1:8080/swagger/
2026-04-25T13:09:04.306Z [INFO] openapi specification is serving at address: http://127.0.0.1:8080/api.json
Feature: action
  Test action-related APIs
2026-04-28T11:50:01.189Z [INFO] initialize start...
...
29 scenarios (27 passed, 2 failed)
256 steps (234 passed, 2 failed, 20 skipped)
34.388410065s
FAIL    github.com/vulcangz/gf2-authz/functional        35.027s
FAIL
```

## Credits

- [eko/authz](https://github.com/eko/authz)
- [GoFrame](https://github.com/gogf/gf)
