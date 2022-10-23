package controller

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

type NewOrder struct {
	CustomerName string        `json:"customer_name" binding:"required"`
	Items        []models.Item `json:"items" binding:"required"`
}

func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}

// GetOrders godoc
// @Summary Get list of all orders
// @ID get-all-orders
// @Description Get all orders
// @Tags order
// @Accept json
// @Produces json
// @Success 200 {array} models.Order
// @Router /orders [get]
func (c Controller) GetAllOrder(ctx *gin.Context) {

	orders := []models.Order{}
	result := c.db.DB.Preload("Items").Find(&orders)
	rows, err := result.Rows()

	if err != nil {
		fmt.Println("Error getting all rows from order table,", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_code": "500",
			"message":    "Error Getting all rows from order table",
		})
	}

	for rows.Next() {
		var order models.Order

		err := rows.Scan(&order)

		if err != nil {
			continue
		}

		orders = append(orders, order)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

// GetOrderByID godoc
// @Summary Get an order data by id
// @ID get-order-by-id
// @Description Get an order data by id
// @Tags orders
// @Accept json
// @Produces json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
func (c Controller) GetOrder(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")

	err := c.db.DB.Preload("Items").Where("id", id).First(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_code": "404",
			"message":    "Order id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, order)

}

// CreateOrder godoc
// @Summary Create a new order
// @ID create-order
// @Description Create a new order with the input payload
// @Tags orders
// @Accept json
// @Produces json
// @Param order body models.Order true "Create Order"
// @Success 201 {object} models.Order
// @Router /orders [post]
func (c Controller) CreateOrder(ctx *gin.Context) {
	var newOrder NewOrder

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "400",
			"message":    err.Error(),
		})
		return
	}

	order := models.Order{CustomerName: newOrder.CustomerName, Items: newOrder.Items, OrderedAt: time.Now()}
	err = c.db.DB.Create(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_code": "500s",
			"message":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created",
		"order":   order,
	})

}

// UpdateOrder godoc
// @Summary Update an order data by id
// @ID update-order-by-id
// @Description Update an order data by id with the input payload
// @Tags orders
// @Accept json
// @Produces json
// @Param id path string true "Order ID"
// @Param newdata body models.Order true "Update Order Data"
// @Success 200 {object} models.Order
// @Router /orders/{id} [put]
func (c Controller) UpdateOrder(ctx *gin.Context) {

	var order models.Order
	id := ctx.Param("id")

	err := c.db.DB.Preload("Items").Where("id", id).First(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_code": "404",
			"message":    "Order id not found",
		})
		return
	}

	order.OrderedAt = time.Now()
	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "400",
			"message":    err.Error(),
		})
		return
	}

	c.db.DB.Save(&order)

	ctx.JSON(http.StatusOK, gin.H{
		"message":        "Update order data successfully",
		"new_order_data": order,
	})
}

// DeleteOrder godoc
// @Summary Delete an order by id
// @ID delete-order-by-id
// @Description Delete an order by id
// @Tags orders
// @Accept json
// @Produces json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id} [delete]
func (c Controller) DeleteOrder(ctx *gin.Context) {
	var order models.Order
	var item models.Item
	id := ctx.Param("id")
	err := c.db.DB.Where("order_id", id).First(&item).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_code": "404",
			"message":    "Order id not found",
		})
		return
	}
	c.db.DB.Delete(&item)

	err = c.db.DB.Where("id", id).First(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_code": "404",
			"message":    "Order id not found",
		})
		return
	}
	c.db.DB.Delete(&order)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %s successfully deleted", id),
	})

}
