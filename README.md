# MallBots monolith demo guide

MallBots is an event-driven retail automation demo built as a modular monolith in Go with an Angular kiosk frontend. The current milestone is a monolith-first kiosk experience backed by real domain services for baskets, customers, stores, payments, ordering, depot, notifications, and search.

This README is the runnable guide for that milestone: how to boot the monolith, how deterministic demo data is prepared, how to use the kiosk flow, and how to verify the happy path.

## What this demo proves

The seeded monolith demo exercises a real customer journey:

- browse a seeded store catalog
- add products to a live basket
- authorize payment and submit checkout
- create an order through the backend saga
- advance fulfillment and payment to final completion in automated proof runs
- observe order state through the live backend search projection

The seeded demo data comes from `internal/demo/spec.go` and is intended to be reproducible.

## Repository layout for the demo

- `cmd/mallbots/main.go` - monolith entrypoint
- `client/soko-bora-web-app/` - Angular kiosk/admin frontend
- `docs/monolith-happy-path.md` - focused proof notes for the seeded happy path
- `testing/e2e/` - seeded end-to-end verification suite
- `docker-compose.yml` - monolith, infra, and observability services
- `docker/.env` - monolith container environment values

## Prerequisites

You need:

- Docker + Docker Compose
- Go 1.24+
- Node.js + npm

## Boot the monolith demo

From the repository root:

```bash
make soko-bora
```

That starts the monolith profile from `docker-compose.yml`, including:

- monolith app on `http://localhost:8080`
- PostgreSQL on `localhost:5432`
- NATS JetStream on `localhost:4222`
- Jaeger UI on `http://localhost:8081`
- Prometheus on `http://localhost:9090`
- Grafana on `http://localhost:3000`
- Pact broker on `http://localhost:9292`

If you prefer the raw compose command:

```bash
docker compose --profile monolith up
```

## Deterministic seed data

The monolith milestone uses deterministic demo entities from `internal/demo/spec.go`.

The seeded happy-path data includes:

- customer: `Demo Shopper`
- stores:
  - `Fresh Harvest Grocers`
  - `Pantry Corner`
- products used in the happy path:
  - `Bananas` x2 at `6`
  - `Rice 2kg` x1 at `18`

Expected seeded basket total:

- `30`

You do not need to hand-seed the database for normal monolith startup. For `@seeded-demo` end-to-end tests, the test suite reseeds the deterministic state automatically when run with `-mono`.

## Start the Angular kiosk

Open a second terminal:

```bash
cd client/soko-bora-web-app
npm install
npm start
```

The Angular dev server runs on `http://localhost:4200` and proxies backend calls to `http://localhost:8080`.

The app routes the main demo surface to `/kiosk`, and that route is protected. If you are not signed in, the frontend redirects you to `/sign-in` first and then returns you to the kiosk flow after authentication.

## Kiosk happy path

The current kiosk milestone supports the following customer-facing flow in the Angular UI:

1. load the seeded demo customer and participating store catalog
2. browse products from the seeded stores
3. add and remove products from the live basket
4. see basket totals update from backend basket state
5. authorize payment and submit checkout
6. refresh live order status from the backend search projection

In other words, the kiosk is no longer a placeholder page. It is backed by real ConnectRPC calls into the monolith.

## Manual demo walkthrough

Once the monolith and Angular app are running:

1. Open `http://localhost:4200`
2. Navigate to the kiosk flow
3. Add seeded products to the basket
   - `Bananas`
   - `Rice 2kg`
4. Confirm the basket total reaches `30`
5. Trigger checkout from the kiosk
6. Watch the order-status panel read live order state from the backend search path

The kiosk demonstrates the customer-side journey. The full cross-service completion proof, including depot progression, invoice payment, notifications, and final search projection, is covered by the automated seeded proof below.

## Automated verification

### Canonical milestone proof

This is the single command that proves the seeded kiosk demo works end to end:

```bash
go test -tags e2e ./testing/e2e/... -run TestEndToEnd -mono -godog.tags=@demo-milestone-proof
```

That focused proof runs the canonical seeded kiosk scenario in:

- `testing/e2e/features/kiosk/shopping.feature`

It verifies the seeded journey through:

- basket start
- seeded item addition
- payment authorization
- checkout
- order creation
- depot shopping list lifecycle
- invoice creation and payment
- notification persistence
- final search projection status

### Seeded demo regression suite

To run the broader seeded-demo scenarios:

```bash
go test -tags e2e ./testing/e2e/... -run TestEndToEnd -mono -godog.tags=@seeded-demo
```

### Frontend verification

From `client/soko-bora-web-app`:

```bash
npm test -- --watch=false --browsers=ChromeHeadless
npm run build
```

## Demo proof narrative

The monolith-first proof covers these backend transitions:

1. the seeded customer starts a basket
2. seeded products are added
3. checkout creates an approved order
4. the create-order saga creates a depot shopping list
5. fulfillment marks the order ready
6. invoice payment completes the order
7. search projects the final backend-visible status

For extra detail, see:

- `docs/monolith-happy-path.md`

## Useful commands

From the repository root:

```bash
make soko-bora
go test ./...
go test -tags e2e ./testing/e2e/... -mono
go generate ./...
```

From `client/soko-bora-web-app`:

```bash
npm start
npm test
npm run build
npm run generate
```

## Stop and clean up

To stop the monolith stack:

```bash
docker compose --profile monolith down
```

To stop the stack and remove volumes with the repo's make target:

```bash
make clean-up-soko-bora
```

If you prefer the raw compose command for the same cleanup:

```bash
docker compose down -v
```

## License

Distributed under the MIT License. See `MIT-LICENSE.txt`.
