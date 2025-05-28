package routes

import (
	"github.com/gin-gonic/gin"
	"pharmacy-backend/controllers"
	"pharmacy-backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/products", controllers.GetProducts)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/products/search", controllers.SearchProducts)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/orders", controllers.CreateOrder)
			protected.GET("/orders", controllers.GetUserOrders)
			protected.PUT("/profile", controllers.UpdateProfile)
			protected.POST("/categories", controllers.AddCategory)
		}
	}
}
