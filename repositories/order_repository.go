package repositories

import (
	"learning-golang/interfaces"
	"learning-golang/models"
)

type orderRepo struct{}

func (r *orderRepo) CreateOrder(o models.Order) (int, error) {
	return o.OrderId, nil
}

func NewOrderRepository() interfaces.OrderRepository {
	return &orderRepo{}
}
