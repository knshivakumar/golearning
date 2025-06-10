package repositories

import (
	"database/sql"
	"learning-golang/interfaces"
	"learning-golang/models"
)

type orderRepo struct {
	db *sql.DB
}

func (r *orderRepo) CreateOrder(o models.Order) (int64, error) {
	result, err := r.db.Exec("insert into orders(orderid, type, price, quanity) values(?, ?, ?, ?)", o.OrderId, o.Type, o.Price, o.Quantity)
	if err != nil {
		return 0, nil
	} else {
		id, _ := result.LastInsertId()
		return id, nil
	}
}

func NewOrderRepository(database *sql.DB) interfaces.OrderRepository {
	return &orderRepo{db: database}
}
