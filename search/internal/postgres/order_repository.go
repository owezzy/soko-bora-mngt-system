package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/stackus/errors"

	"github.com/owezzy/soko-bora-mngt-system/internal/postgres"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/application"
	"github.com/owezzy/soko-bora-mngt-system/search/internal/models"
)

type OrderRepository struct {
	tableName string
	db        postgres.DB
}

var _ application.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(tableName string, db postgres.DB) OrderRepository {
	return OrderRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r OrderRepository) Add(ctx context.Context, order *models.Order) error {
	const query = `INSERT INTO %s (
order_id, customer_id, customer_name,
items, status, product_ids, store_ids,
created_at) VALUES (
$1, $2, $3,
$4, $5, $6, $7,
$8)`

	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	productIDs := make(IDArray, len(order.Items))
	storeMap := make(map[string]struct{})
	for i, item := range order.Items {
		productIDs[i] = item.ProductID
		storeMap[item.StoreID] = struct{}{}
	}
	storeIDs := make(IDArray, 0, len(storeMap))
	for storeID, _ := range storeMap {
		storeIDs = append(storeIDs, storeID)
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		order.OrderID, order.CustomerID, order.CustomerName,
		items, order.Status, productIDs, storeIDs,
		order.CreatedAt,
	)
	return err
}

func (r OrderRepository) UpdateStatus(ctx context.Context, orderID, status string) error {
	const query = `UPDATE %s SET status = $2 WHERE order_id = $1`

	_, err := r.db.ExecContext(ctx, r.table(query), orderID, status)
	return err
}

func (r OrderRepository) Search(ctx context.Context, search application.SearchOrders) ([]*models.Order, error) {
	const selectQuery = `SELECT order_id, customer_id, customer_name, items, status, created_at FROM %s`

	args := make([]any, 0, 10)
	clauses := make([]string, 0, 8)

	if search.Filters.CustomerID != "" {
		args = append(args, search.Filters.CustomerID)
		clauses = append(clauses, fmt.Sprintf("customer_id = $%d", len(args)))
	}
	if !search.Filters.After.IsZero() {
		args = append(args, search.Filters.After)
		clauses = append(clauses, fmt.Sprintf("created_at >= $%d", len(args)))
	}
	if !search.Filters.Before.IsZero() {
		args = append(args, search.Filters.Before)
		clauses = append(clauses, fmt.Sprintf("created_at <= $%d", len(args)))
	}
	if len(search.Filters.StoreIDs) > 0 {
		args = append(args, IDArray(search.Filters.StoreIDs))
		clauses = append(clauses, fmt.Sprintf("store_ids && $%d", len(args)))
	}
	if len(search.Filters.ProductIDs) > 0 {
		args = append(args, IDArray(search.Filters.ProductIDs))
		clauses = append(clauses, fmt.Sprintf("product_ids && $%d", len(args)))
	}
	if search.Filters.Status != "" {
		args = append(args, search.Filters.Status)
		clauses = append(clauses, fmt.Sprintf("status = $%d", len(args)))
	}
	if search.Filters.MinTotal > 0 {
		args = append(args, search.Filters.MinTotal)
		clauses = append(clauses, fmt.Sprintf("%s >= $%d", orderTotalExpr(), len(args)))
	}
	if search.Filters.MaxTotal > 0 {
		args = append(args, search.Filters.MaxTotal)
		clauses = append(clauses, fmt.Sprintf("%s <= $%d", orderTotalExpr(), len(args)))
	}

	query := r.table(selectQuery)
	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	query += " ORDER BY created_at DESC, order_id DESC"

	offset, err := parseOffset(search.Next)
	if err != nil {
		return nil, err
	}

	args = append(args, search.Limit)
	query += fmt.Sprintf(" LIMIT $%d", len(args))
	args = append(args, offset)
	query += fmt.Sprintf(" OFFSET $%d", len(args))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			err = errors.Wrap(closeErr, "closing order rows")
		}
	}(rows)

	orders := make([]*models.Order, 0, search.Limit)
	for rows.Next() {
		order, scanErr := scanOrder(rows)
		if scanErr != nil {
			return nil, scanErr
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r OrderRepository) Get(ctx context.Context, orderID string) (*models.Order, error) {
	const query = `SELECT customer_id, customer_name, items, status, created_at FROM %s WHERE order_id = $1`

	order := &models.Order{
		OrderID: orderID,
	}

	var itemData []byte
	err := r.db.QueryRowContext(ctx, r.table(query), orderID).Scan(&order.CustomerID, &order.CustomerName, &itemData, &order.Status, &order.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrNotFound.Msgf("order with id: `%s` does not exist", orderID)
		}
		return nil, err
	}

	var items []models.Item
	err = json.Unmarshal(itemData, &items)
	if err != nil {
		return nil, err
	}
	order.Items = items
	order.Total = calculateTotal(items)

	return order, nil
}

func (r OrderRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}

type IDArray []string

func (a *IDArray) Scan(src any) error {
	var sep = []byte(",")

	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return errors.ErrInvalidArgument.Msgf("IDArray: unsupported type: %T", src)
	}

	ids := make([]string, bytes.Count(data, sep))
	for i, id := range bytes.Split(bytes.Trim(data, "{}"), sep) {
		ids[i] = string(id)
	}

	*a = ids

	return nil
}

func (a IDArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	if len(a) == 0 {
		return "{}", nil
	}
	// unsafe way to do this; assumption is all ids are UUIDs
	return fmt.Sprintf("{%s}", strings.Join(a, ",")), nil
}

func parseOffset(next string) (int, error) {
	if next == "" {
		return 0, nil
	}

	offset, err := strconv.Atoi(next)
	if err != nil || offset < 0 {
		return 0, errors.ErrInvalidArgument.Msg("search next cursor must be a non-negative offset")
	}

	return offset, nil
}

func orderTotalExpr() string {
	return `COALESCE((SELECT SUM((item->>'Price')::double precision * (item->>'Quantity')::integer) FROM jsonb_array_elements(convert_from(items, 'UTF8')::jsonb) item), 0)`
}

func scanOrder(scanner interface{ Scan(dest ...any) error }) (*models.Order, error) {
	order := &models.Order{}
	var itemData []byte
	if err := scanner.Scan(&order.OrderID, &order.CustomerID, &order.CustomerName, &itemData, &order.Status, &order.CreatedAt); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(itemData, &order.Items); err != nil {
		return nil, err
	}
	order.Total = calculateTotal(order.Items)

	return order, nil
}

func calculateTotal(items []models.Item) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}
