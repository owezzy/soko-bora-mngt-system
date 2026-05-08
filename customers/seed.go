package customers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/owezzy/soko-bora-mngt-system/customers/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/customers/internal/constants"
	"github.com/owezzy/soko-bora-mngt-system/customers/internal/domain"
	"github.com/owezzy/soko-bora-mngt-system/internal/demo"
	"github.com/owezzy/soko-bora-mngt-system/internal/di"
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

	spec := demo.Spec().Customer
	repo := di.Get(ctx, constants.CustomersRepoKey).(domain.CustomerRepository)
	app := di.Get(ctx, constants.ApplicationKey).(application.App)

	customer, findErr := repo.Find(ctx, spec.ID)
	if findErr != nil {
		if errors.Is(findErr, sql.ErrNoRows) {
			err = app.RegisterCustomer(ctx, application.RegisterCustomer{
				ID:        spec.ID,
				Name:      spec.Name,
				SmsNumber: spec.SMSNumber,
			})
			return err
		}
		err = findErr
		return err
	}

	if customer.Name != spec.Name || customer.SmsNumber != spec.SMSNumber {
		err = repo.Update(ctx, &domain.Customer{
			Aggregate: domain.NewCustomer(spec.ID).Aggregate,
			Name:      spec.Name,
			SmsNumber: spec.SMSNumber,
			Enabled:   true,
		})
		if err != nil {
			return err
		}
	}

	if !customer.Enabled {
		err = app.EnableCustomer(ctx, application.EnableCustomer{ID: spec.ID})
	}

	return err
}
