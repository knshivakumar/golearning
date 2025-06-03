package routers

import (
	"learning-golang/controllers"
	"learning-golang/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, orderController *controllers.OrderController) *gin.Engine {
	r := gin.Default()

	//user endpoint
	r.POST("/register", userController.CreateUser)
	r.POST("/login", userController.Login)

	//Protected endpoints
	auth := r.Group("/auth")
	//intercept the JWT token and if it valid grant access to the below routes
	auth.Use(middleware.JWTAuthMiddleware())
	r.GET("/getusers", userController.GetAllUsers)

	r.POST("/createOrder", orderController.CreateOrder)

	return r
}
