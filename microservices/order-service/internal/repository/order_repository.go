package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/hexhoc/order-service/internal/entity"
	"github.com/hexhoc/order-service/pkg/datasource/postgres"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type OrderInterface interface {
	FindAll(ctx context.Context, withItems bool, limit uint32, offset uint32) ([]*entity.Order, error)
	FindById(ctx context.Context, withItems bool, id string) (*entity.Order, error)
	Save(ctx context.Context, product *entity.Order) (string, error)
	Update(ctx context.Context, id string, product *entity.Order) error
	Delete(ctx context.Context, id string) error
}

type OrderRepository struct {
	db *postgres.Postgres
}

// TODO: Добавить транзакции БД
func NewOrderRepository(db *postgres.Postgres) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) FindAll(ctx context.Context, withItems bool, limit uint32, offset uint32) ([]*entity.Order, error) {

	query := `
	SELECT 
		id,
		customer_id,
		address,
		status,
		is_deleted,
		created_at,
		updated_at
	FROM orders
	LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[string]*entity.Order)
	var orderIds []string
	for rows.Next() {
		order := r.orderRowMapper(rows)
		ordersMap[order.Id] = order
		orderIds = append(orderIds, order.Id)
	}

	if withItems {
		orderItems, err := r.getAllOrderItem(ctx, orderIds)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		for _, orderItem := range orderItems {
			ordersMap[orderItem.OrderId].OrderItems = append(ordersMap[orderItem.OrderId].OrderItems, orderItem)
		}
	}

	var ordersList []*entity.Order
	for _, order := range ordersMap {
		ordersList = append(ordersList, order)
	}
	return ordersList, nil
}

func (r *OrderRepository) FindById(ctx context.Context, withItems bool, id string) (*entity.Order, error) {
	query := `
	SELECT 
		id,
		customer_id,
		address,
		status,
		is_deleted,
		created_at,
		updated_at
	FROM orders
	WHERE id = $1
	`

	rows, err := r.db.Pool.Query(ctx, query, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var order *entity.Order
	for rows.Next() {
		order = r.orderRowMapper(rows)
	}

	if withItems {
		orderItems, err := r.getAllOrderItem(ctx, []string{order.Id})
		if err != nil {
			log.Error(err)
			return nil, err
		}
		order.OrderItems = orderItems
	}

	return order, nil
}

func (r *OrderRepository) Save(ctx context.Context, order *entity.Order) (string, error) {
	//TODO: add transaction
	query := `INSERT INTO orders(customer_id, address, status, is_deleted, created_at, updated_at) 
			  VALUES ($1,$2,$3,$4,$5,$6)
			  RETURNING id`

	rows, err := r.db.Pool.Query(
		ctx, query,
		order.CustomerId,
		order.Address,
		order.Status,
		order.IsDeleted,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Printf("failed create order %s", err.Error())
		return "", fmt.Errorf("failed create order: %w", err)
	}

	var orderId string
	rows.Next()
	rows.Scan(&orderId)
	rows.Close()

	err = r.insertOrderItem(ctx, orderId, order.OrderItems)
	if err != nil {
		log.Printf("failed create order %s", err.Error())
		return "", fmt.Errorf("failed create order: %w", err)
	}

	return orderId, nil
}

func (r *OrderRepository) Update(ctx context.Context, id string, order *entity.Order) error {

	query := `
	UPDATE orders SET 
		customer_id = $1,
		address = $2,
		status = $3,
		is_deleted = $4,
		updated_at = $5
	WHERE orders.id = $6
	`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		order.CustomerId,
		order.Address,
		order.Status,
		order.IsDeleted,
		time.Now(),
		id,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	// delete order item
	err = r.deleteOrderItem(ctx, order.OrderItems)
	if err != nil {
		return err
	}
	err = r.insertOrderItem(ctx, order.Id, order.OrderItems)
	if err != nil {
		return err
	}
	err = r.updateOrderItem(ctx, order.OrderItems)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Order update row affected %d", ct.RowsAffected()))

	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders WHERE orders.id = $1`
	ct, err := r.db.Pool.Exec(
		ctx, query,
		id,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Orders delete row affected %d", ct.RowsAffected()))

	return nil
}

func (r *OrderRepository) insertOrderItem(ctx context.Context, orderId string, orderItems []*entity.OrderItem) error {

	// Batch insert order items
	queryItems := `INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`
	batch := &pgx.Batch{}
	for _, item := range orderItems {
		if item.Id != "" {
			continue
		}
		price, _ := item.Price.Float64()
		batch.Queue(queryItems, orderId, item.ProductId, item.Quantity, price)
	}

	if batch.Len() == 0 {
		// Nothing to insert
		return nil
	}

	results := r.db.Pool.SendBatch(ctx, batch)
	defer results.Close()

	for _, v := range orderItems {
		_, err := results.Exec()
		if err != nil {
			log.Printf("order item %s already exists", v.Id)
			return fmt.Errorf("unable to insert row: %w", err)
		}
	}

	return results.Close()
}

func (r *OrderRepository) updateOrderItem(ctx context.Context, orderItems []*entity.OrderItem) error {

	// Batch insert order items
	query := `
	UPDATE order_items SET
		order_id = $1, 
		product_id = $2, 
		quantity = $3, 
		price = $4 
	WHERE order_items.id = $5`

	batch := &pgx.Batch{}
	for _, item := range orderItems {
		if item.Id == "" {
			continue
		}
		price, _ := item.Price.Float64()
		batch.Queue(query, item.OrderId, item.ProductId, item.Quantity, price, item.Id)
	}

	// nothing to update
	if batch.Len() == 0 {
		return nil
	}

	results := r.db.Pool.SendBatch(ctx, batch)
	defer results.Close()

	for _, v := range orderItems {
		_, err := results.Exec()
		if err != nil {
			log.Printf("order item %s already exists", v.Id)
			return fmt.Errorf("unable to insert row: %w", err)
		}

	}

	return results.Close()
}

func (r *OrderRepository) deleteOrderItem(ctx context.Context, orderItems []*entity.OrderItem) error {

	var idsList []string
	for _, orderItem := range orderItems {
		if orderItem.Id == "" {
			continue
		}
		idsList = append(idsList, orderItem.Id)
	}

	if len(idsList) == 0 {
		return nil
	}

	idsString := "{" + strings.Join(idsList, ",") + "}"

	// Batch insert order items
	query := `DELETE FROM order_items WHERE order_items.id <> ANY($5)`
	ct, err := r.db.Pool.Exec(
		ctx, query,
		idsString,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Orders delete row affected %d", ct.RowsAffected()))

	return nil
}

func (r *OrderRepository) getAllOrderItem(ctx context.Context, orderIds []string) ([]*entity.OrderItem, error) {

	query := `
	SELECT 
		order_items.id         as item_id,
		order_items.product_id as product_id,
		order_items.quantity   as item_quantity,
		order_items.price      as item_price
	FROM order_items
	WHERE order_items.order_id = ANY($1)
`

	idsString := "{" + strings.Join(orderIds, ",") + "}"
	rows, err := r.db.Pool.Query(ctx, query, idsString)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var orderItems []*entity.OrderItem
	for rows.Next() {
		orderItems = append(orderItems, r.itemRowMapper(rows))
	}

	return orderItems, nil
}

func (r *OrderRepository) orderRowMapper(rows pgx.Rows) *entity.Order {
	var item entity.Order
	err := rows.Scan(
		&item.Id,
		&item.CustomerId,
		&item.Address,
		&item.Status,
		&item.IsDeleted,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		log.Error(fmt.Errorf("error while iterating dataset %w", err))
	}

	return &item
}

func (r *OrderRepository) itemRowMapper(rows pgx.Rows) *entity.OrderItem {
	var item entity.OrderItem
	err := rows.Scan(
		&item.Id,
		&item.OrderId,
		&item.ProductId,
		&item.Quantity,
		&item.Price,
	)

	if err != nil {
		log.Error(fmt.Errorf("error while iterating dataset %w", err))
	}

	return &item
}
