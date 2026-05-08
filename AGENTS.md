# AGENTS.md

## Project

Event-driven modular monolith (Go 1.24) following the "Event-Driven Architecture in Golang" (Packt) pattern. Can run as a single monolith or as separate microservices via Docker Compose profiles.

Module: `github.com/owezzy/soko-bora-mngt-system`

Top-level `README.md` is still template boilerplate. Treat `go.mod`, `Makefile`, `docker-compose.yml`, `buf.yaml`, module `generate.go`, and code under `cmd/` / `internal/` as the source of truth.

## Architecture

### Domain Modules (bounded contexts)

`baskets/`, `cosec/`, `customers/`, `depot/`, `notifications/`, `ordering/`, `payments/`, `search/`, `stores/`

Each module follows the same layout (except `cosec/` — see below):
```
<module>/
  module.go          # Startup() wires DI container, registers adapters
  generate.go        # go:generate directives (buf, mockery, swagger)
  buf.gen.yaml       # Per-module protobuf codegen config
  internal/
    constants/       # String-based DI keys (ServiceName, repo keys, handler keys)
    application/     # Use cases / command-query handlers
    domain/          # Aggregates, events, repositories (interfaces)
    grpc/            # gRPC server (driver adapter)
    rest/            # REST gateway + swagger
    handlers/        # Domain event & integration event handlers
    postgres/        # Repository implementations (driven adapter)
  <module>pb/        # Generated protobuf Go code
  <module>client/    # Generated swagger client
  migrations/        # Per-module SQL migrations (embedded)
  cmd/               # Microservice entrypoint (for split deployment)
```

**Module exceptions:**
- `cosec/` — Saga orchestrator. No protobuf, no gRPC/REST, no `generate.go`. Has `internal/saga.go` + `internal/models/` instead. Uses `internal/sec/` (saga engine).
- `notifications/` — Smaller module. Still has protobuf generation, persistence/cache code, inbox handling, and a gRPC server, but no REST gateway and no per-request DI container.
- `baskets/` — Has `ui/` subdirectory (embedded web assets for kiosk UI).

### Shared Infrastructure (`internal/`)

- `am/` - Async messaging abstractions
- `ddd/` - DDD building blocks (aggregates, events, entities)
- `di/` - Dependency injection container
- `es/` - Event sourcing (aggregate store, snapshots)
- `jetstream/` - NATS JetStream stream adapter
- `tm/` - Transaction manager, outbox processor
- `postgres/` - Shared PG helpers (event store, outbox, inbox, snapshots)
- `registry/` - Type registry for serialization/deserialization
- `system/` - System bootstrap (DB, NATS, gRPC, HTTP mux, config, waiter)
- `monolith/` - Monolith runner
- `config/` - App config via `envconfig`
- `web/` - Embedded web UI assets
- `rpc/` - gRPC server setup

### Entrypoints

- **Monolith**: `cmd/mallbots/main.go` — registers all modules, runs migrations from `migrations/` (embedded SQL via goose), starts web + RPC + stream
- **Microservices**: Each module has `<module>/cmd/` for independent deployment
- **Busywork client**: `cmd/busywork/` — test/load client

## Key Patterns

- **Event Sourcing**: Aggregates persisted via event store + snapshots (`internal/es/`)
- **Outbox Pattern**: Transactional outbox for reliable messaging (`internal/tm/`). Table names are module-scoped (`<module>.outbox`, `<module>.inbox`).
- **NATS JetStream**: Message broker for inter-module async communication
- **DI Container**: `internal/di/` scoped containers, constants-based keys per module
- **Protobuf + ConnectRPC + gRPC-Gateway**: Each module defines protos, generates Go gRPC, ConnectRPC, REST gateway, OpenAPI spec, and TypeScript client stubs

## Commands

```bash
# Run monolith (Docker)
make soko-bora                    # docker compose --profile monolith up (needs Docker)
docker compose --profile microservices up  # Run as microservices

# Code generation (protobuf, mocks, swagger clients)
make generate                     # go generate ./...

# Install protoc plugins
make install-tools

# Protobuf linting
buf lint

# Tests
go test ./...                              # Unit + contract tests
go test -tags e2e ./testing/e2e/...        # E2E (requires running app + infra)
go test -tags e2e ./testing/e2e/... -mono  # E2E against monolith DB
```

## Testing

- **Unit tests**: Standard `_test.go` files in module internals (e.g. `baskets/internal/domain/basket_test.go`)
- **Contract tests**: Pact-based (`pact-go/v2`) in `*_contract_test.go`, build tag `contract`. Pact broker at `localhost:9292`
- **Integration tests**: `*_integration_test.go`, build tag `integration`. Use testcontainers-go for Postgres + NATS. Docker daemon must be running.
- **E2E tests**: `testing/e2e/` — Cucumber/Godog BDD, build tag `e2e`, features in `testing/e2e/features/`. Connects to `localhost:8080`

```bash
# Run by category
go test ./...                                        # Unit only (no build tags)
go test -tags integration ./...                      # Integration (needs Docker)
go test -tags contract ./...                         # Contract (needs pact broker)
go test -tags e2e ./testing/e2e/... -mono            # E2E (needs running app)
```

Integration tests use `testify/suite` with `SetupSuite`/`TearDownSuite` lifecycle for container management.

## Infrastructure (Docker Compose)

Services: postgres (12-alpine), nats (JetStream), otel-collector, jaeger, prometheus, grafana, pact-broker

- Env: `docker/.env` (loaded by Makefile and compose)
- DB init scripts: `docker/database/`
- Each microservice gets its own PG database/user/schema (e.g. `stores` DB, `stores_user`, `stores` schema)
- Monolith uses single `mallbots` DB

Ports: HTTP `8080`, gRPC `8085`, Jaeger UI `8081`, Grafana `3000`, Prometheus `9090`, Pact `9292`

## Code Generation

`go generate ./...` is not uniform across packages:
1. Most domain modules run `buf generate`
2. Only `ordering/`, `payments/`, and `stores/` also run `mockery` + `swagger generate client`
3. Shared packages like `internal/am`, `internal/ddd`, and `internal/es` generate mocks too
4. Frontend TypeScript stubs are emitted into `client/soko-bora-web-app/src/proto/` from module `buf.gen.yaml` files

**Do not hand-edit** files in `*pb/`, `*client/`, or `client/soko-bora-web-app/src/proto/`.

## Frontend

Angular app at `client/soko-bora-web-app/`. See `client/soko-bora-web-app/AGENTS.md` for ConnectRPC, proxy, and Fuse-specific frontend rules.

## Env & Config

Config loaded via `kelseyhightower/envconfig`. Key env vars: `ENVIRONMENT`, `PG_CONN`, `NATS_URL`, `OTEL_SERVICE_NAME`, `OTEL_EXPORTER_OTLP_ENDPOINT`, `RPC_SERVICES`.

Env file for Docker: `docker/.env`. Uses `stackus/dotenv` for local `.envrc` loading.

## Logging

`rs/zerolog` throughout. Errors via `stackus/errors` (wraps standard errors with gRPC status codes).

## Architecture Decisions

Documented in `docs/ADL/`. `0002-use-a-modular-monolith-architecture.md` is the key repo-specific design note and explicitly relies on `/internal/` packages to preserve module autonomy.

<!-- BEGIN BEADS INTEGRATION v:1 profile:minimal hash:ca08a54f -->
## Beads Issue Tracker

This project uses **bd (beads)** for issue tracking. Run `bd prime` to see full workflow context and commands.

### Quick Reference

```bash
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --claim  # Claim work
bd close <id>         # Complete work
```

### Rules

- Use `bd` for ALL task tracking — do NOT use TodoWrite, TaskCreate, or markdown TODO lists
- Run `bd prime` for detailed command reference and session close protocol
- Use `bd remember` for persistent knowledge — do NOT use MEMORY.md files

## Session Completion

**When ending a work session**, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **File issues for remaining work** - Create issues for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   git pull --rebase
   bd dolt push
   git push
   git status  # MUST show "up to date with origin"
   ```
5. **Clean up** - Clear stashes, prune remote branches
6. **Verify** - All changes committed AND pushed
7. **Hand off** - Provide context for next session

**CRITICAL RULES:**
- Work is NOT complete until `git push` succeeds
- NEVER stop before pushing - that leaves work stranded locally
- NEVER say "ready to push when you are" - YOU must push
- If push fails, resolve and retry until it succeeds
<!-- END BEADS INTEGRATION -->
