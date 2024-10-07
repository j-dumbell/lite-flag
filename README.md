# lite-flag
![example workflow](https://github.com/j-dumbell/lite-flag/actions/workflows/test-build.yml/badge.svg?branch=main)

Lightweight feature flag service.

### Stack
- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) for generating API handlers and types from the Open API spec.
- [chi](https://github.com/go-chi/chi) for routing and middleware.
- [net/http](https://pkg.go.dev/net/http) for the HTTP server.
- [database/sql](https://pkg.go.dev/database/sql) for the SQL client.


## Development
Ensure [Docker](https://www.docker.com/) is installed.

#### Bootstrap the DB
```bash
make boostrap
```

#### Start the API
```bash
make start-api
```

#### Run all tests
```bash
make test
```

#### Regenerate the API handlers and types
```bash
make oapi-gen
```

#### Regenerate the clients
```bash
make gen-clients
```
