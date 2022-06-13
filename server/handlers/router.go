package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	services *services.Services
}

func NewRouter(services *services.Services) *Router {
	return &Router{services: services}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", r.signUp)
		auth.POST("/sign-in", r.signIn)
	}

	api := router.Group("/api")
	{
		api.GET("/cab-man", r.getAllCabMans)
		api.GET("/cab-man/:id", r.getOneCabMan)
		api.GET("/orders", r.getOrdersByStatus)
		//api.GET("/orders/driver=:driver_id/off=:offset", r.getOrdersByDriverId)
		api.GET("test", r.test)
		api.POST("/cab-man", r.createCabMan)
		api.PUT("/cab-man/:id", r.updateCabMan)
		api.DELETE("/cab-man/:id", r.deleteCabMan)
	}

	return router
}
