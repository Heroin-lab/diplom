package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Router struct {
	services *services.Services
}

func NewRouter(services *services.Services) *Router {
	return &Router{services: services}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", r.signUp)
		auth.POST("/sign-in", r.signIn)
	}

	api := router.Group("/api")
	{
		api.GET("/offenses", r.getAllUserOffenses)
		api.POST("/offenses", r.createUserOffense)
		api.DELETE("/offenses", r.deleteUserOffenses)
		api.PUT("/offenses", r.updateUserOffense)
		//api.GET("/cab-man", r.getAllCabMans)
		//api.GET("/cab-man/:id", r.getOneCabMan)
		//api.GET("/orders", r.getOrdersByStatus)
		//api.GET("/orders/driver=:driver_id/off=:offset", r.getOrdersByDriverId)
		//api.GET("test", r.test)
		//api.POST("/cab-man", r.createCabMan)
		//api.PUT("/cab-man/:id", r.updateCabMan)
		//api.DELETE("/cab-man/:id", r.deleteCabMan)
	}

	return router
}
