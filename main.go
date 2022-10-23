package main

import (
	"assignment-2/controller"
	"assignment-2/database"
	"fmt"

	_ "assignment-2/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Orders API
// @version 1.0
// @description This is a service for managing orders
// @contact.name API Support
// @contact.email suryaaasaputra.s@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	db, err := database.StartDB()

	if err != nil {
		fmt.Println("error start database", err)
		return
	}
	ctl := controller.New(db)

	r := gin.Default()

	r.GET("/orders", ctl.GetAllOrder)
	r.GET("/orders/:id", ctl.GetOrder)
	r.POST("/orders", ctl.CreateOrder)
	r.PUT("/orders/:id", ctl.UpdateOrder)
	r.DELETE("/orders/:id", ctl.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("localhost:8080")
}
