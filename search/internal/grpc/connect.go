package grpc

import (
	"context"
	"database/sql"
	"net/http"

	"connectrpc.com/connect"

	"github.com/owezzy/soko-bora-mngt-system/internal/di"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/constants"
	"github.com/owezzy/soko-bora-mngt-system/search/searchpb"
	"github.com/owezzy/soko-bora-mngt-system/search/searchpb/searchpbconnect"
)

type connectServer struct {
	app application.Application
	searchpbconnect.UnimplementedSearchServiceHandler
}

type connectServerTx struct {
	c di.Container
	searchpbconnect.UnimplementedSearchServiceHandler
}

func NewConnectHandlerTx(container di.Container) (string, http.Handler) {
	return searchpbconnect.NewSearchServiceHandler(connectServerTx{c: container})
}

func (s connectServer) SearchOrders(ctx context.Context, request *connect.Request[searchpb.SearchOrdersRequest]) (*connect.Response[searchpb.SearchOrdersResponse], error) {
	resp, err := server{app: s.app}.SearchOrders(ctx, request.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s connectServer) GetOrder(ctx context.Context, request *connect.Request[searchpb.GetOrderRequest]) (*connect.Response[searchpb.GetOrderResponse], error) {
	resp, err := server{app: s.app}.GetOrder(ctx, request.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s connectServer) GetDemoBootstrap(ctx context.Context, request *connect.Request[searchpb.GetDemoBootstrapRequest]) (*connect.Response[searchpb.GetDemoBootstrapResponse], error) {
	resp, err := server{app: s.app}.GetDemoBootstrap(ctx, request.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s connectServerTx) SearchOrders(ctx context.Context, request *connect.Request[searchpb.SearchOrdersRequest]) (resp *connect.Response[searchpb.SearchOrdersResponse], err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	return connectServer{app: di.Get(ctx, constants.ApplicationKey).(application.Application)}.SearchOrders(ctx, request)
}

func (s connectServerTx) GetOrder(ctx context.Context, request *connect.Request[searchpb.GetOrderRequest]) (resp *connect.Response[searchpb.GetOrderResponse], err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	return connectServer{app: di.Get(ctx, constants.ApplicationKey).(application.Application)}.GetOrder(ctx, request)
}

func (s connectServerTx) GetDemoBootstrap(ctx context.Context, request *connect.Request[searchpb.GetDemoBootstrapRequest]) (resp *connect.Response[searchpb.GetDemoBootstrapResponse], err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	return connectServer{app: di.Get(ctx, constants.ApplicationKey).(application.Application)}.GetDemoBootstrap(ctx, request)
}

func closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	}
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
