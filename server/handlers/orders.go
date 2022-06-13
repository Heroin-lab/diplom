package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/server/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (r *Router) getOrdersByStatus(c *gin.Context) {
	statusId := c.DefaultQuery("status_id", "0")
	offsetNum := c.DefaultQuery("offset", "0")

	ordersList, err := r.services.OrdersServices.GetOrdersByStatus(statusId, offsetNum)
	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.OrdersResponse{
		Data: ordersList,
	})
}

func (r *Router) getOrdersByDriverId(c *gin.Context) {
	driverId := c.Query("driver_id")
	offsetNum := c.DefaultQuery("offset", "0")

	orderInfo, err := r.services.OrdersServices.GetOrdersByDriverId(driverId, offsetNum)
	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.OrdersResponse{
		Data: orderInfo,
	})
}

func (r *Router) test(c *gin.Context) {
	startLocations := []string{"54.234236235", "55.234236235", "56.234236235"}
	endLocations := []string{"57.234236235", "58.234236235", "59.234236235"}
	err := r.services.OrdersServices.CreateOrders(startLocations, endLocations)
	if err != nil {
		logrus.Error(err)
		return
	}
}
