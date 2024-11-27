package database

import (
	"database/sql"

	"github.com/lauronicolas/curso-go/Clean-Arch-APP/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List() ([]*entity.Order, error) {
	var orders []*entity.Order
	rows, err := r.Db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order entity.Order
		if err = rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
