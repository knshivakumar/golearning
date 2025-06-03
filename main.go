package main

import (
	"database/sql"
	"learning-golang/configs"
	"learning-golang/dependency"
	"learning-golang/routers"
)

func main() {
	//Get DB connection
	var db *sql.DB
	defer db.Close()
	db = configs.ConnectDB()

	//Build DI Container
	container := dependency.BuildContainer(db)

	//Setup router

	r := routers.SetupRouter(container.UserController, container.OrderController)

	//Start server
	r.Run(":4000")

}
