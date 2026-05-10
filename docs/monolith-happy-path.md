# Monolith happy path

This backend proof uses the seeded demo bootstrap and exercises the monolith path across:

- baskets
- ordering
- depot
- payments
- notifications
- search

## Seeded demo data

The monolith startup seeds deterministic demo data from `internal/demo/spec.go`.

- Customer: `Demo Shopper`
- Store: `Fresh Harvest Grocers`
- Store: `Pantry Corner`
- Products used in the happy path:
  - `Bananas` x2 at `6`
  - `Rice 2kg` x1 at `18`

The scenario total is `30`.

## Run the proof

Start the monolith and infrastructure first, then run:

```bash
go test -tags e2e ./testing/e2e/... -mono
```

For the single seeded kiosk proof that now serves as the minimum milestone check, run:

```bash
go test -tags e2e ./testing/e2e/... -run TestEndToEnd -mono -godog.tags=@demo-milestone-proof
```

## Canonical seeded kiosk proof

The scenario `Seeded kiosk happy path completes end to end` in `testing/e2e/features/kiosk/shopping.feature` is the canonical proof for the monolith-first kiosk milestone.

It is the one scenario intended to prove that the seeded kiosk demo is reproducible from startup state through final backend completion.

## What the proof covers

The seeded kiosk proof verifies that:

1. a seeded demo customer can start a basket
2. seeded demo products can be added and checked out
3. checkout creates an approved order
4. the create-order saga creates a depot shopping list
5. the shopping list can be assigned and completed
6. completing the shopping list readies the order and creates an invoice
7. paying the invoice completes the order
8. search projects `Ready For Pickup` and then `Completed`
9. notifications persist observable lifecycle rows for the order
