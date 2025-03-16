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
TODO...

## Credits

- [eko/authz](https://github.com/eko/authz)
- [GoFrame](https://github.com/gogf/gf)
