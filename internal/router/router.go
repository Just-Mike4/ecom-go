package router

import (
	"net/http"

	"github.com/Just-Mike4/ecom-go/internal/carts"
	"github.com/Just-Mike4/ecom-go/internal/orders"
	"github.com/Just-Mike4/ecom-go/internal/payments"
	"github.com/Just-Mike4/ecom-go/internal/products"
	"github.com/Just-Mike4/ecom-go/internal/users"
	"github.com/go-chi/chi/v5"
)

func Setup() http.Handler {
	r := chi.NewRouter()

	// Initialize handlers
	userHandler := users.NewHandler()
	productHandler := products.NewHandler()
	cartHandler := carts.NewHandler()
	orderHandler := orders.NewHandler()
	paymentHandler := payments.NewHandler()

	// API v1 routes
	r.Route("/api/v1", func(r chi.Router) {
		// Authentication endpoints
		r.Post("/auth/register", userHandler.Register)
		r.Post("/auth/login", userHandler.Login)
		r.Post("/auth/logout", userHandler.Logout)
		r.Post("/auth/refresh", userHandler.RefreshToken)
		r.Post("/auth/password/forgot", userHandler.PasswordForgot)
		r.Post("/auth/password/reset", userHandler.PasswordReset)
		r.Post("/auth/verify-email", userHandler.VerifyEmail)
		r.Get("/auth/me", userHandler.AuthMe)

		// Admin authentication
		r.Post("/auth/admin/register", userHandler.AdminRegister)
		r.Post("/auth/admin/login", userHandler.AdminLogin)

		// Vendor authentication
		r.Post("/auth/vendor/register", userHandler.VendorRegister)
		r.Post("/auth/vendor/login", userHandler.VendorLogin)

		// User management
		r.Patch("/users/me", userHandler.UpdateUserMe)
		r.Post("/users/me/avatar", userHandler.UploadAvatar)
		r.Patch("/users/me/password", userHandler.ChangePassword)
		r.Get("/users", userHandler.ListUsers)
		r.Get("/users/{id}", userHandler.GetUserByID)
		r.Delete("/users/{id}", userHandler.DeleteUser)

		// Addresses
		r.Get("/users/me/addresses", userHandler.ListAddresses)
		r.Post("/users/me/addresses", userHandler.CreateAddress)
		r.Patch("/users/me/addresses/{id}", userHandler.UpdateAddress)
		r.Delete("/users/me/addresses/{id}", userHandler.DeleteAddress)
		r.Patch("/users/me/addresses/{id}/default", userHandler.SetDefaultAddress)

		// Notifications
		r.Get("/notifications", userHandler.ListNotifications)
		r.Patch("/notifications/{id}/read", userHandler.MarkNotificationRead)
		r.Post("/notifications/read-all", userHandler.MarkAllNotificationsRead)

		// Admin endpoints
		r.Get("/admin/dashboard", userHandler.AdminDashboard)
		r.Get("/admin/analytics", userHandler.AdminAnalytics)
		r.Get("/admin/users", userHandler.AdminListUsers)
		r.Post("/admin/users/{id}/ban", userHandler.AdminBanUser)
		r.Post("/admin/users/{id}/unban", userHandler.AdminUnbanUser)
		r.Get("/admin/coupons", userHandler.AdminListCoupons)
		r.Post("/admin/coupons", userHandler.AdminCreateCoupon)
		r.Get("/admin/coupons/validate/{code}", userHandler.AdminValidateCoupon)

		// Products
		r.Get("/products", productHandler.ListProducts)
		r.Get("/products/search", productHandler.SearchProducts)
		r.Get("/products/{id}", productHandler.GetProduct)
		r.Post("/products/admin/{id}/images", productHandler.AdminUploadProductImage)
		r.Delete("/products/admin/{id}/images/{image_id}", productHandler.AdminDeleteProductImage)
		r.Get("/products/admin/{id}/inventory", productHandler.AdminGetProductInventory)
		r.Patch("/products/admin/{id}/inventory", productHandler.AdminUpdateProductInventory)

		// Categories
		r.Get("/categories", productHandler.ListCategories)
		r.Post("/categories/admin", productHandler.AdminCreateCategory)
		r.Get("/categories/{id}", productHandler.GetCategory)
		r.Patch("/categories/{id}", productHandler.UpdateCategory)
		r.Delete("/categories/{id}", productHandler.DeleteCategory)
		r.Get("/categories/{id}/products", productHandler.GetCategoryProducts)

		// Reviews
		r.Get("/products/{id}/reviews", productHandler.ListProductReviews)
		r.Post("/products/{id}/reviews", productHandler.CreateProductReview)
		r.Patch("/reviews/{id}", productHandler.UpdateReview)
		r.Delete("/reviews/{id}", productHandler.DeleteReview)
		r.Get("/users/me/reviews", productHandler.ListUserReviews)

		// Vendor products
		r.Get("/vendor/products", productHandler.VendorListProducts)
		r.Post("/vendor/products", productHandler.VendorCreateProduct)
		r.Get("/vendor/products/{id}", productHandler.VendorGetProduct)
		r.Patch("/vendor/products/{id}", productHandler.VendorUpdateProduct)
		r.Delete("/vendor/products/{id}", productHandler.VendorDeleteProduct)

		// Cart
		r.Get("/cart", cartHandler.GetCart)
		r.Post("/cart/items", cartHandler.AddCartItem)
		r.Patch("/cart/items/{item_id}", cartHandler.UpdateCartItem)
		r.Delete("/cart/items/{item_id}", cartHandler.DeleteCartItem)
		r.Delete("/cart/clear", cartHandler.ClearCart)
		r.Post("/cart/sync", cartHandler.SyncCart)

		// Wishlist
		r.Get("/wishlist", cartHandler.GetWishlist)
		r.Post("/wishlist/items", cartHandler.AddWishlistItem)
		r.Delete("/wishlist/items/{product_id}", cartHandler.DeleteWishlistItem)

		// Orders
		r.Post("/orders", orderHandler.CreateOrder)
		r.Get("/orders", orderHandler.ListOrders)
		r.Get("/orders/{id}", orderHandler.GetOrder)
		r.Get("/orders/{id}/tracking", orderHandler.GetOrderTracking)
		r.Post("/orders/{id}/cancel", orderHandler.CancelOrder)
		r.Post("/orders/{id}/return-request", orderHandler.CreateReturnRequest)

		// Admin orders
		r.Get("/admin/orders", orderHandler.AdminListOrders)
		r.Patch("/admin/orders/{id}/status", orderHandler.AdminUpdateOrderStatus)
		r.Post("/admin/orders/{id}/ship", orderHandler.AdminShipOrder)
		r.Post("/admin/orders/{id}/refund", orderHandler.AdminRefundOrder)

		// Shipping
		r.Get("/shipping/rates", orderHandler.GetShippingRates)
		r.Get("/shipping/methods", orderHandler.GetShippingMethods)

		// Vendor orders
		r.Get("/vendor/orders", orderHandler.VendorListOrders)
		r.Patch("/vendor/orders/{id}/status", orderHandler.VendorUpdateOrderStatus)

		// Payments
		r.Get("/payments", paymentHandler.ListPayments)
		r.Post("/payments/initialize", paymentHandler.InitializePayment)
		r.Get("/payments/verify/{reference}", paymentHandler.VerifyPayment)
		r.Post("/payments/webhook", paymentHandler.PaymentWebhook)
		r.Get("/payments/{id}", paymentHandler.GetPayment)
		r.Post("/payments/{id}/refund", paymentHandler.RefundPayment)
	})

	return r
}
