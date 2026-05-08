package stores

import (
	"context"
	"database/sql"
	stderrors "errors"
	"fmt"

	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/internal/demo"
	"github.com/owezzy/soko-bora-mngt-system/internal/di"
	"github.com/owezzy/soko-bora-mngt-system/stores/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/stores/internal/application/commands"
	"github.com/owezzy/soko-bora-mngt-system/stores/internal/application/queries"
	"github.com/owezzy/soko-bora-mngt-system/stores/internal/constants"
	"github.com/owezzy/soko-bora-mngt-system/stores/internal/domain"
)

func seedDemoData(ctx context.Context, container di.Container) (err error) {
	ctx = container.Scoped(ctx)
	tx := di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	app := di.Get(ctx, constants.ApplicationKey).(application.App)
	mall := di.Get(ctx, constants.MallRepoKey).(domain.MallRepository)
	catalog := di.Get(ctx, constants.CatalogRepoKey).(domain.CatalogRepository)

	for _, storeSpec := range demo.Spec().Stores {
		if err = ensureStore(ctx, app, mall, storeSpec); err != nil {
			return err
		}
		for _, productSpec := range storeSpec.Products {
			if err = ensureProduct(ctx, app, catalog, storeSpec, productSpec); err != nil {
				return err
			}
		}
	}

	return nil
}

func ensureStore(ctx context.Context, app application.App, mall domain.MallRepository, spec demo.StoreSpec) error {
	store, err := app.GetStore(ctx, queries.GetStore{ID: spec.ID})
	if err != nil {
		if !isMissing(err) {
			return err
		}
		if err = app.CreateStore(ctx, commands.CreateStore{
			ID:       spec.ID,
			Name:     spec.Name,
			Location: spec.Location,
		}); err != nil {
			return err
		}
		store, err = mall.Find(ctx, spec.ID)
		if err != nil {
			return err
		}
	}

	if store.Name != spec.Name {
		if err = app.RebrandStore(ctx, commands.RebrandStore{ID: spec.ID, Name: spec.Name}); err != nil {
			return err
		}
		store.Name = spec.Name
	}

	if store.Location != spec.Location {
		return fmt.Errorf("demo store %s exists with location %q, expected %q", spec.ID, store.Location, spec.Location)
	}

	if spec.Participating && !store.Participating {
		return app.EnableParticipation(ctx, commands.EnableParticipation{ID: spec.ID})
	}
	if !spec.Participating && store.Participating {
		return app.DisableParticipation(ctx, commands.DisableParticipation{ID: spec.ID})
	}

	return nil
}

func ensureProduct(ctx context.Context, app application.App, catalog domain.CatalogRepository, storeSpec demo.StoreSpec, spec demo.ProductSpec) error {
	product, err := catalog.Find(ctx, spec.ID)
	if err != nil {
		if !isMissing(err) {
			return err
		}
		if err = app.AddProduct(ctx, commands.AddProduct{
			ID:          spec.ID,
			StoreID:     storeSpec.ID,
			Name:        spec.Name,
			Description: spec.Description,
			SKU:         spec.SKU,
			Price:       spec.Price,
		}); err != nil {
			return err
		}
		product, err = catalog.Find(ctx, spec.ID)
		if err != nil {
			return err
		}
	}

	if product.StoreID != storeSpec.ID {
		return fmt.Errorf("demo product %s belongs to store %s, expected %s", spec.ID, product.StoreID, storeSpec.ID)
	}
	if product.SKU != spec.SKU {
		return fmt.Errorf("demo product %s exists with sku %q, expected %q", spec.ID, product.SKU, spec.SKU)
	}

	if product.Name != spec.Name || product.Description != spec.Description {
		if err = app.RebrandProduct(ctx, commands.RebrandProduct{
			ID:          spec.ID,
			Name:        spec.Name,
			Description: spec.Description,
		}); err != nil {
			return err
		}
	}

	if product.Price < spec.Price {
		return app.IncreaseProductPrice(ctx, commands.IncreaseProductPrice{ID: spec.ID, Price: spec.Price})
	}
	if product.Price > spec.Price {
		return app.DecreaseProductPrice(ctx, commands.DecreaseProductPrice{ID: spec.ID, Price: spec.Price})
	}

	return nil
}

func isMissing(err error) bool {
	return stderrors.Is(err, sql.ErrNoRows) || stderrors.Is(err, errors.ErrNotFound)
}
