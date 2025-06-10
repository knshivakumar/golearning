package services

import (
	"learning-golang/interfaces"
	"learning-golang/models"
)

type OrderService interface {
	CreateOrder(models.Order) (orderId int64, err error)
}

type orderService struct {
	orepo interfaces.OrderRepository
}

func NewOrderService(r interfaces.OrderRepository) OrderService {
	return &orderService{orepo: r}
}

func (s *orderService) CreateOrder(o models.Order) (int64, error) {
	return s.orepo.CreateOrder(o)
}
