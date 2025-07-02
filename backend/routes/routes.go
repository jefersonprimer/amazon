package routes

import (
	"backend/handlers"
	"backend/middlewares" // Vamos ajustar este em breve para o JWT

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rotas Públicas (não exigem autenticação)
	public := router.Group("/api/v1")
	{
		public.POST("/register", handlers.RegisterUser)
		public.POST("/login", handlers.LoginUser)
		public.POST("/forgot-password", handlers.ForgotPasswordHandler) // Nova rota
		public.POST("/reset-password", handlers.ResetPassword)          // Nova rota

		// Exemplo de rotas públicas para listagens
		public.GET("/products", handlers.GetProducts)
		public.GET("/products/:id", handlers.GetProductByID)
		public.GET("/categories", handlers.GetCategories)
		public.GET("/vendors", handlers.GetVendors)

		// Rotas públicas de review
		public.GET("/reviews", handlers.GetReviews)
		public.GET("/reviews/:id", handlers.GetReviewByID)

		// Rotas públicas de imagens de produto
		public.GET("/product-images", handlers.GetProductImages)
		public.GET("/product-images/:id", handlers.GetProductImageByID)
	}

	// Rotas Protegidas (exigem autenticação JWT)
	protected := router.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware()) // Aplica o middleware JWT para este grupo
	{
		// Rotas de Usuário (precisarão de handlers específicos, mas já protegidas)
		// protected.GET("/users/:id", handlers.GetUserByID)
		// protected.PUT("/users/:id", handlers.UpdateUser)

		// Exemplo de rotas protegidas
		// protected.POST("/products", handlers.CreateProduct)
		// protected.PUT("/products/:id", handlers.UpdateProduct)
		// protected.DELETE("/products/:id", handlers.DeleteProduct)

		// Rotas de Vendedor (com middleware de autorização de papel)
		vendorRoutes := protected.Group("/vendor")
		vendorRoutes.Use(middlewares.VendorAuthMiddleware())
		{
			// Exemplo: vendorRoutes.GET("/my-products", handlers.GetVendorProducts)
		}

		// Rotas de Admin (com middleware de autorização de papel)
		adminRoutes := protected.Group("/admin")
		adminRoutes.Use(middlewares.AdminAuthMiddleware())
		{
			// Exemplo: adminRoutes.GET("/users", handlers.GetUsers)
			// Exemplo: adminRoutes.PUT("/users/:id/status", handlers.UpdateUserStatus)
		}

		// Rotas de review protegidas (apenas usuários autenticados podem criar)
		protected.POST("/reviews", handlers.CreateReview)
		// Se quiser proteger update/delete, adicione:
		// protected.PUT("/reviews/:id", handlers.UpdateReview)
		// protected.DELETE("/reviews/:id", handlers.DeleteReview)

		// Rotas protegidas de imagens de produto
		protected.POST("/product-images", handlers.CreateProductImage)
		protected.PUT("/product-images/:id", handlers.UpdateProductImage)
		protected.DELETE("/product-images/:id", handlers.DeleteProductImage)

		// Rotas protegidas de pedidos (Order)
		protected.POST("/orders", handlers.CreateOrder)
		protected.GET("/orders/:id", handlers.GetOrderByID)
		protected.GET("/orders", handlers.GetOrdersByUser)
		protected.PUT("/orders/:id/status", handlers.UpdateOrderStatus)
		protected.DELETE("/orders/:id", handlers.DeleteOrder)
	}
}
