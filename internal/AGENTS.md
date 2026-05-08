# internal/ — Shared Infrastructure

Shared packages used by all domain modules. Do not import domain module code here.

`internal/` is where cross-cutting building blocks live. Domain logic stays in module-local `*/internal/` packages, not here.

## DI Container (`di/`)

String-keyed, scoped container. Not type-safe — keys defined as string constants in each module's `internal/constants/constants.go`.

- `AddSingleton(key, factory)` — shared across all requests
- `AddScoped(key, factory)` — per-request/per-transaction, resolved via `di.Get(ctx, key)`
- Cyclic dependency detection via `tracked.go`
- Every module creates its own container in `module.go` → `Root()`

## Event Sourcing (`es/`)

- `AggregateRepository[T]` — generic repo for event-sourced aggregates
- `AggregateStore` with middleware chain: event store → snapshot store
- Snapshots: versioned (e.g. `StoreV1`, `ProductV1`), configurable strategy (TODO in code)
- All aggregate/event types must be registered in the module's `registrations()` function via `registry/serdes`

## Async Messaging (`am/`, `jetstream/`)

- `am.EventPublisher` / `am.MessagePublisher` — publish domain/integration events
- `am.MessageSubscriber` — subscribe to NATS JetStream subjects
- Outbox pattern: messages go to PG outbox table first, then `tm.OutboxProcessor` sends to NATS
- Each module has its own outbox/inbox tables (e.g. `stores.outbox`, `stores.inbox`)
- `amotel/` — OpenTelemetry context propagation for async messaging
- `amprom/` — Prometheus counters for sent/received messages

## Handler Registration Pattern

Handlers are split across two files per concern:
- `handlers.go` — pure handler logic (domain event handlers, integration event handlers)
- `*_transaction.go` — wraps handlers with DI scoped context for transaction management

The `*_transaction.go` files use `di.Get(ctx, key)` to resolve scoped dependencies (DB transaction, repos) per request.

## Saga Engine (`sec/`)

Used only by `cosec/` module. Defines saga steps, compensation, and orchestration. See `cosec/internal/saga.go` for the saga definition.

## System Bootstrap (`system/`)

`system.Service` interface provides: `DB()`, `JS()` (JetStream), `RPC()` (gRPC), `Mux()` (HTTP), `Config()`, `Logger()`, `Waiter()`.

All modules receive `system.Service` in their `Startup(ctx, svc)` method.

`system.WaitForWeb()` wraps the chi mux in h2c + CORS, exposes `/liveness` and `/metrics`, and is the real HTTP entrypoint for both REST and Connect handlers.

## Key Packages Not to Confuse

- `postgres/` (here) — shared PG helpers (event store, outbox, inbox, snapshots)
- `postgresotel/` — tracing wrapper around DB/tx usage
- `<module>/internal/postgres/` — module-specific repository implementations
- `am/` — messaging abstractions (interfaces)
- `jetstream/` — NATS JetStream concrete implementation of `am/` interfaces
- `errorsotel/` — OpenTelemetry helpers around error handling
- `logger/` — zerolog initialization from app config
- `rpc/` — shared gRPC setup helpers
- `waiter/` — lifecycle/shutdown coordination used by `system/`
