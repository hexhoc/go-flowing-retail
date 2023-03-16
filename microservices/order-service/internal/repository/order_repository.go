package repository

import (
	"fmt"
	"math/big"
	"time"

	"github.com/hexhoc/order-service/internal/entity"
	"github.com/hexhoc/order-service/pkg/datasource/postgres"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type OrderInterface interface {
	FindAll(ctx context.Context, limit uint32, offset uint32) ([]*entity.Order, error)
	FindById(ctx context.Context, id string) (*entity.Order, error)
	Save(ctx context.Context, product *entity.Order) error
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

func (r *OrderRepository) FindAll(ctx context.Context, limit uint32, offset uint32) ([]*entity.Order, error) {

	query := `
	SELECT 
		id,
		customer_id,
		address,
		status,
		is_deleted,
		created_at,
		update_at,
	FROM orders as o 
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

	queryItems := `
	SELECT 
		order_items.id         as item_id,
		order_items.product_id as product_id,
		order_items.quantity   as item_quantity,
		order_items.price      as item_price
	FROM order_items
	WHERE order_items.order_id IN ($1)
	`
	rowsItems, err := r.db.Pool.Query(ctx, queryItems, orderIds)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rowsItems.Close()

	for rowsItems.Next() {
		orderItem := r.itemRowMapper(rowsItems)
		ordersMap[orderItem.OrderId].OrderItems = append(ordersMap[orderItem.OrderId].OrderItems, orderItem)
	}

	var ordersList []*entity.Order
	for _, order := range ordersMap {
		ordersList = append(ordersList, order)
	}
	return ordersList, nil
}

func (r *OrderRepository) FindById(ctx context.Context, id string) (*entity.Order, error) {
	query := `
	SELECT 
		orders.id,
		orders.customer_id,
		orders.address,
		orders.status,
		orders.is_deleted,
		orders.created_at,
		orders.update_at,
		order_items.id         as item_id,
		order_items.product_id as product_id,
		order_items.quantity   as item_quantity,
		order_items.price      as item_price
	FROM orders
	LEFT JOIN order_items ON orders.id = order_items.order_id
	WHERE orders.id = $1
	`

	rows, err := r.db.Pool.Query(ctx, query, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var item *entity.Order
	orderMap := make(map[string]*entity.Order)
	for rows.Next() {
		r.orderItemRowMapper(rows, orderMap)
	}

	for _, v := range orderMap {
		item = v
	}

	return item, nil
}

func (r *OrderRepository) Save(ctx context.Context, item *entity.Order) error {
	query := `
	INSERT INTO orders(customer_id, address, status, is_deleted, created_at, updated_at) 
	VALUES ($1,$2,$3,$4,$5)
	`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		item.CustomerId,
		item.Address,
		item.Status,
		item.IsDeleted,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Order update row affected %d", ct.RowsAffected()))

	// Batch insert order items

	queryItems := `INSERT INTO order_items (order_id, product_id, quantity, price) 
	VALUES (@order_id, @product_id, @quantity, @price)`
	batch := &pgx.Batch{}
	for _, v := range item.OrderItems {
		batch.Queue(queryItems, v.OrderId, v.ProductId, v.Quantity, v.Price)
	}

	results := r.db.Pool.SendBatch(ctx, batch)
	defer results.Close()

	for _, v := range item.OrderItems {
		_, err := results.Exec()
		if err != nil {
			log.Printf("order item %s already exists", v.Id)
			return fmt.Errorf("unable to insert row: %w", err)
		}
	}

	return results.Close()
}

func (r *OrderRepository) Update(ctx context.Context, id string, item *entity.Order) error {
	query := `
	UPDATE orders SET 
		customer_id = $1,
		address = $2,
		status = $3,
		is_deleted = $4,
		update_at = $5,
	WHERE orders.id = $6
	`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		item.CustomerId,
		item.Address,
		item.Status,
		item.IsDeleted,
		time.Now(),
		id,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	// Batch insert order items
	queryItems := `
	UPDATE order_items SET
		order_id = $1, 
		product_id = $2, 
		quantity = $3, 
		price = $4 
	WHERE order_items.id = $5`

	batch := &pgx.Batch{}
	for _, v := range item.OrderItems {
		batch.Queue(queryItems, v.OrderId, v.ProductId, v.Quantity, v.Price, v.Id)
	}

	results := r.db.Pool.SendBatch(ctx, batch)
	defer results.Close()

	for _, v := range item.OrderItems {
		_, err := results.Exec()
		if err != nil {
			log.Printf("order item %s already exists", v.Id)
			return fmt.Errorf("unable to insert row: %w", err)
		}

	}
	log.Info(fmt.Sprintf("Order update row affected %d", ct.RowsAffected()))

	return results.Close()
}

func (r *OrderRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders where orders.id = $1`
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

func (r *OrderRepository) orderItemRowMapper(rows pgx.Rows, orderMap map[string]*entity.Order) {

	var (
		orderId       string
		customerId    uint32
		address       string
		status        uint32
		isDeleted     bool
		createdAt     time.Time
		updatedAt     time.Time
		itemId        string
		itemProductId uint32
		itemQuantity  uint32
		itemPrice     big.Float
	)
	rows.Scan(
		&orderId,
		&customerId,
		&address,
		&status,
		&isDeleted,
		&createdAt,
		&updatedAt,
		&itemId,
		&itemProductId,
		&itemQuantity,
		&itemPrice,
	)

	if _, ok := orderMap[orderId]; !ok {
		order := entity.Order{
			Id:         orderId,
			CustomerId: customerId,
			Address:    address,
			Status:     status,
			IsDeleted:  isDeleted,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			OrderItems: []*entity.OrderItem{},
		}

		orderMap[orderId] = &order
	}

	orderItem := entity.OrderItem{
		Id:        itemId,
		OrderId:   orderId,
		ProductId: itemProductId,
		Quantity:  itemQuantity,
		Price:     itemPrice,
	}

	orderMap[orderId].OrderItems = append(orderMap[orderId].OrderItems, &orderItem)
}
