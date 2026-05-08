package grpc

import (
	"context"
	"strconv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/owezzy/soko-bora-mngt-system/internal/errorsotel"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
	"github.com/owezzy/soko-bora-mngt-system/search/searchpb"
)

type server struct {
	app application.Application
	searchpb.UnimplementedSearchServiceServer
}

func RegisterServer(_ context.Context, app application.Application, registrar grpc.ServiceRegistrar) error {
	searchpb.RegisterSearchServiceServer(registrar, server{app: app})
	return nil
}

func (s server) SearchOrders(ctx context.Context, request *searchpb.SearchOrdersRequest) (*searchpb.SearchOrdersResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.Int("Limit", int(request.GetLimit())),
		attribute.String("Next", request.GetNext()),
	)

	orders, err := s.app.SearchOrders(ctx, application.SearchOrders{
		Filters: searchFiltersFromProto(request.GetFilters()),
		Next:    request.GetNext(),
		Limit:   int(request.GetLimit()),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	protoOrders := make([]*searchpb.Order, len(orders))
	for i, order := range orders {
		protoOrders[i] = orderFromDomain(order)
	}

	return &searchpb.SearchOrdersResponse{
		Orders: protoOrders,
		Next:   nextOffset(request, len(orders)),
	}, nil
}

func (s server) GetOrder(ctx context.Context, request *searchpb.GetOrderRequest) (*searchpb.GetOrderResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("OrderID", request.GetId()))

	order, err := s.app.GetOrder(ctx, application.GetOrder{OrderID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &searchpb.GetOrderResponse{Order: orderFromDomain(order)}, nil
}

func (s server) GetDemoBootstrap(ctx context.Context, request *searchpb.GetDemoBootstrapRequest) (*searchpb.GetDemoBootstrapResponse, error) {
	span := trace.SpanFromContext(ctx)
	bootstrap, err := s.app.GetDemoBootstrap(ctx)
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return demoBootstrapFromDomain(bootstrap), nil
}

func searchFiltersFromProto(filters *searchpb.SearchOrdersRequest_Filters) application.Filters {
	if filters == nil {
		return application.Filters{}
	}

	searchFilters := application.Filters{
		CustomerID: filters.GetCustomerId(),
		StoreIDs:   append([]string(nil), filters.GetStoreIds()...),
		ProductIDs: append([]string(nil), filters.GetProductIds()...),
		MinTotal:   filters.GetMinTotal(),
		MaxTotal:   filters.GetMaxTotal(),
		Status:     filters.GetStatus(),
	}

	if after := filters.GetAfter(); after != nil {
		searchFilters.After = after.AsTime()
	}
	if before := filters.GetBefore(); before != nil {
		searchFilters.Before = before.AsTime()
	}

	return searchFilters
}

func orderFromDomain(order *models.Order) *searchpb.Order {
	items := make([]*searchpb.Order_Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = itemFromDomain(item)
	}

	return &searchpb.Order{
		OrderId:      order.OrderID,
		CustomerId:   order.CustomerID,
		CustomerName: order.CustomerName,
		Items:        items,
		Total:        order.Total,
		Status:       order.Status,
	}
}

func itemFromDomain(item models.Item) *searchpb.Order_Item {
	return &searchpb.Order_Item{
		ProductId:   item.ProductID,
		StoreId:     item.StoreID,
		ProductName: item.ProductName,
		StoreName:   item.StoreName,
		Price:       item.Price,
		Quantity:    int64(item.Quantity),
	}
}

func demoBootstrapFromDomain(bootstrap application.DemoBootstrap) *searchpb.GetDemoBootstrapResponse {
	stores := make([]*searchpb.DemoStore, 0, len(bootstrap.Stores))
	products := make([]*searchpb.DemoProduct, 0)
	for _, store := range bootstrap.Stores {
		stores = append(stores, &searchpb.DemoStore{
			Id:   store.ID,
			Name: store.Name,
		})
		for _, product := range store.Products {
			products = append(products, &searchpb.DemoProduct{
				Id:        product.ID,
				StoreId:   store.ID,
				StoreName: store.Name,
				Name:      product.Name,
			})
		}
	}

	return &searchpb.GetDemoBootstrapResponse{
		Customer: &searchpb.DemoCustomer{
			Id:   bootstrap.Customer.ID,
			Name: bootstrap.Customer.Name,
		},
		Stores:   stores,
		Products: products,
	}
}

func nextOffset(request *searchpb.SearchOrdersRequest, returned int) string {
	limit := int(request.GetLimit())
	if limit <= 0 {
		limit = application.DefaultSearchLimit
	}
	if returned < limit {
		return ""
	}

	offset, err := strconv.Atoi(request.GetNext())
	if request.GetNext() == "" {
		err = nil
		offset = 0
	}
	if err != nil || offset < 0 {
		return ""
	}

	return strconv.Itoa(offset + returned)
}
