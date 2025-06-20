package interfaces

import (
	"learning-golang/models"
)

type OrderRepository interface {
	CreateOrder(models.Order) (orderId int64, err error)
}
