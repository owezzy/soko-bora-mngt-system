//go:build e2e

package e2e

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/messages-go/v16"
	"github.com/cucumber/godog"
	"github.com/go-openapi/runtime/client"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rdumont/assistdog"
	"github.com/stackus/errors"
)

var useMonoDB = flag.Bool("mono", false, "Use mono DB resources")
var godogTags = flag.String("godog.tags", "", "Filter scenarios by tag expression")

var assist = assistdog.NewDefault()

type lastResponseKey struct{}
type lastErrorKey struct{}

type feature interface {
	init(cfg featureConfig) error
	register(ctx *godog.ScenarioContext)
	reset()
}

type featureConfig struct {
	transport *client.Runtime
	useMonoDB bool
}

func TestEndToEnd(t *testing.T) {
	assist.RegisterComparer(float64(0.0), func(raw string, actual interface{}) error {
		af, ok := actual.(float64)
		if !ok {
			return fmt.Errorf("%v is not a float64", actual)
		}
		ef, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return err
		}

		if ef != af {
			return fmt.Errorf("expected %v, but got %v", ef, af)
		}

		return nil
	})
	assist.RegisterParser(float64(0.0), func(raw string) (interface{}, error) {
		return strconv.ParseFloat(raw, 64)
	})

	cfg := featureConfig{
		transport: client.New("localhost:8080", "/", nil),
		useMonoDB: *useMonoDB,
	}

	features, err := func(fs ...feature) ([]feature, error) {
		features := make([]feature, len(fs))
		for i, f := range fs {
			err := f.init(cfg)
			if err != nil {
				return features, err
			}
			features[i] = f
		}
		return features, nil
	}(
		&basketsFeature{},
		&customersFeature{},
		&notificationsFeature{},
		&ordersFeature{},
		&paymentsFeature{},
		&storesFeature{},
	)
	if err != nil {
		t.Fatal(err)
	}

	featurePaths := []string{
		"features/baskets",
		"features/customers",
		"features/kiosk",
		"features/orders",
		"features/stores",
	}

	suite := godog.TestSuite{
		Name: "mallbots-e2e",
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			ctx.Step(`^I receive a "([^"]*)" error$`, iReceiveAError)
			ctx.Step(`^(?:ensure |expect )?the returned error message is "([^"]*)"$`, iReceiveAError)
			for _, f := range features {
				f.register(ctx)
			}
			ctx.Before(func(ctx context.Context, s *godog.Scenario) (context.Context, error) {
				seededDemo := scenarioHasTag(s, "@seeded-demo")
				for _, f := range features {
					f.reset()
				}
				if seededDemo && cfg.useMonoDB {
					if err := reseedSeededDemoMonolithState(ctx); err != nil {
						return ctx, err
					}
				}

				return setSeededDemoScenario(ctx, seededDemo), nil
			})
		},
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     featurePaths,
			Randomize: -1,
			Tags:      *godogTags,
		},
	}

	if status := suite.Run(); status != 0 {
		t.Error("end to end feature test failed with status:", status)
	}
}

func scenarioHasTag(s *messages.Pickle, tag string) bool {
	for _, pickleTag := range s.Tags {
		if pickleTag.Name == tag {
			return true
		}
	}

	return false
}

func iReceiveAError(ctx context.Context, msg string) error {
	err := lastError(ctx)
	if err == nil {
		return errors.Wrap(errors.ErrUnknown, "expected error to not be nil")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "Message:"+msg) && !strings.Contains(errMsg, msg) {
		return errors.Wrapf(errors.ErrInvalidArgument, "expected: %s: got: %s", msg, err.Error())
	}

	return nil
}

func setLastResponseAndError(ctx context.Context, resp any, err error) context.Context {
	return context.WithValue(
		context.WithValue(ctx, lastResponseKey{}, resp),
		lastErrorKey{}, err,
	)
}

func lastResponseWas(ctx context.Context, resp any) error {
	r := ctx.Value(lastResponseKey{})
	if reflect.ValueOf(r).Kind() == reflect.Ptr && reflect.ValueOf(r).IsNil() {
		e := ctx.Value(lastErrorKey{})
		if e == nil {
			return errors.ErrUnknown.Msg("no last response or error")
		}
		return e.(error)
	}
	if reflect.TypeOf(r) == reflect.TypeOf(resp) {
		return nil
	}
	return errors.ErrBadRequest.Msgf("last request was `%v`", r)
}

func lastResponse(ctx context.Context) any {
	return ctx.Value(lastResponseKey{})
}

func lastError(ctx context.Context) error {
	e := ctx.Value(lastErrorKey{})
	if e == nil {
		return nil
	}
	return e.(error)
}
