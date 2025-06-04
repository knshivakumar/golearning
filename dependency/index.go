package dependency

import (
	"database/sql"
	"learning-golang/controllers"
	"learning-golang/repositories"
	"learning-golang/services"
)

type Container struct {
	UserController  *controllers.UserController
	OrderController *controllers.OrderController
}

func BuildContainer(db *sql.DB) *Container {

	//wiring dependency manually
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	return &Container{
		UserController:  userController,
		OrderController: orderController,
	}
}
