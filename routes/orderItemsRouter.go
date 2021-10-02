package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderitems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderitems/:orderItem_id", controllers.GetOrderItemByOrder())
	incomingRoutes.GET("/orderitems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.POST("/orderitems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderitems/:orderItem_id", controllers.UpdateOrderItem())
}
