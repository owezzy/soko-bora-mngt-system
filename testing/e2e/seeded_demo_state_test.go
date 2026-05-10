//go:build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/owezzy/soko-bora-mngt-system/internal/demo"
)

const monoConnString = "postgres://mallbots_user:mallbots_pass@localhost:5432/mallbots?sslmode=disable"

func reseedSeededDemoMonolithState(ctx context.Context) error {
	db, err := sql.Open("pgx", monoConnString)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	spec := demo.Spec()

	if err = upsertSeededDemoCustomer(ctx, tx, spec.Customer); err != nil {
		return err
	}

	for _, store := range spec.Stores {
		if err = upsertSeededDemoStore(ctx, tx, store); err != nil {
			return err
		}
		for _, product := range store.Products {
			if err = upsertSeededDemoProduct(ctx, tx, store, product); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

func upsertSeededDemoCustomer(ctx context.Context, tx *sql.Tx, customer demo.CustomerSpec) error {
	query := `
		INSERT INTO customers.customers (id, name, sms_number, enabled)
		VALUES ($1, $2, $3, TRUE)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name,
		    sms_number = EXCLUDED.sms_number,
		    enabled = EXCLUDED.enabled
	`
	if _, err := tx.ExecContext(ctx, query, customer.ID, customer.Name, customer.SMSNumber); err != nil {
		return err
	}

	query = `
		INSERT INTO notifications.customers_cache (id, name, sms_number)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name,
		    sms_number = EXCLUDED.sms_number
	`
	if _, err := tx.ExecContext(ctx, query, customer.ID, customer.Name, customer.SMSNumber); err != nil {
		return err
	}

	query = `
		INSERT INTO search.customers_cache (id, name)
		VALUES ($1, $2)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name
	`
	_, err := tx.ExecContext(ctx, query, customer.ID, customer.Name)
	return err
}

func upsertSeededDemoStore(ctx context.Context, tx *sql.Tx, store demo.StoreSpec) error {
	query := `
		INSERT INTO stores.stores (id, name, location, participating)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name,
		    location = EXCLUDED.location,
		    participating = EXCLUDED.participating
	`
	if _, err := tx.ExecContext(ctx, query, store.ID, store.Name, store.Location, store.Participating); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO baskets.stores_cache (id, name)
		VALUES ($1, $2)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name
	`, store.ID, store.Name); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO depot.stores_cache (id, name, location)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name,
		    location = EXCLUDED.location
	`, store.ID, store.Name, store.Location); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `
		INSERT INTO search.stores_cache (id, name)
		VALUES ($1, $2)
		ON CONFLICT (id) DO UPDATE
		SET name = EXCLUDED.name
	`, store.ID, store.Name); err != nil {
		return err
	}

	return nil
}

func upsertSeededDemoProduct(ctx context.Context, tx *sql.Tx, store demo.StoreSpec, product demo.ProductSpec) error {
	query := `
		INSERT INTO stores.products (id, store_id, name, description, sku, price)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE
		SET store_id = EXCLUDED.store_id,
		    name = EXCLUDED.name,
		    description = EXCLUDED.description,
		    sku = EXCLUDED.sku,
		    price = EXCLUDED.price
	`
	if _, err := tx.ExecContext(ctx, query, product.ID, store.ID, product.Name, product.Description, product.SKU, product.Price); err != nil {
		return err
	}

	type cacheStatement struct {
		query string
		args  []any
	}

	stmts := []cacheStatement{
		{
			query: `
				INSERT INTO baskets.products_cache (id, store_id, name, price)
				VALUES ($1, $2, $3, $4)
				ON CONFLICT (id) DO UPDATE
				SET store_id = EXCLUDED.store_id,
				    name = EXCLUDED.name,
				    price = EXCLUDED.price
			`,
			args: []any{product.ID, store.ID, product.Name, product.Price},
		},
		{
			query: `
				INSERT INTO depot.products_cache (id, store_id, name)
				VALUES ($1, $2, $3)
				ON CONFLICT (id) DO UPDATE
				SET store_id = EXCLUDED.store_id,
				    name = EXCLUDED.name
			`,
			args: []any{product.ID, store.ID, product.Name},
		},
		{
			query: `
				INSERT INTO search.products_cache (id, store_id, name)
				VALUES ($1, $2, $3)
				ON CONFLICT (id) DO UPDATE
				SET store_id = EXCLUDED.store_id,
				    name = EXCLUDED.name
			`,
			args: []any{product.ID, store.ID, product.Name},
		},
	}

	for _, stmt := range stmts {
		if _, err := tx.ExecContext(ctx, stmt.query, stmt.args...); err != nil {
			return fmt.Errorf("upserting seeded demo product %s cache: %w", product.ID, err)
		}
	}

	return nil
}
