package main

import (
	"foodlive/component"
	"foodlive/middleware"
	"foodlive/modules/authsso/fbsso/fbssotransport/ginfbsso"
	"foodlive/modules/authsso/googlesso/googlessotransport/gingooglesso"
	"foodlive/modules/cart/carttransport/gincart"
	"foodlive/modules/category/categorytransport/gincategory"
	"foodlive/modules/food/foodtransport/ginfood"
	"foodlive/modules/order/ordertransport/ginorder"
	"foodlive/modules/restaurant/restauranttransport/ginrestaurant"
	"foodlive/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"foodlive/modules/restaurantowner/restaurantownertransport/ginrestaurantowner"
	"foodlive/modules/restaurantrating/restaurantratingtransport/ginrestaurantrating"
	"foodlive/modules/statistic/statistictransport/ginstatistic"
	"foodlive/modules/upload/uploadtransport/ginupload"
	"foodlive/modules/user/usertransport/ginuser"
	"foodlive/modules/useraddress/useraddresstransport/ginuseraddress"
	"foodlive/modules/userdevicetoken/userdevicetokentransport/ginuserdevicetoken"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx component.AppContext) {
	r.Use(middleware.Recover(appCtx))
	v1Route(r, appCtx)
}

func v1Route(r *gin.Engine, appCtx component.AppContext) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", ginuser.UserReigster(appCtx))
		v1.POST("/activate", ginuser.UserActiveAccount(appCtx))
		v1.POST("/resend-otp-active", ginuser.ResendOTPActive(appCtx))
		v1.POST("/forgot-password", ginuser.UserForgotPassword(appCtx))
		v1.POST("/reset-password", ginuser.UserResetPassword(appCtx))

		v1.POST("/login", ginuser.UserLogin(appCtx))
		v1.POST("/login/fb", ginfbsso.UserFacebookLogin(appCtx))
		v1.POST("/login/gg", gingooglesso.UserGoogleLogin(appCtx))
		v1.PUT("/sso/update", middleware.RequireSSOAuth(appCtx), ginuser.AccountSSOUpdate(appCtx))

		v1.POST("/admin-login", ginuser.AdminLogin(appCtx))

		v1.POST("/upload", ginupload.Upload(appCtx))
		v1.POST("/ipn", ginorder.HandleWebHookPayment(appCtx))

		v1.GET("/order/crypto/:order_id", ginorder.FindOrderCryptoInWeb(appCtx))

		restaurant := v1.Group("/restaurant", middleware.RequireAuth(appCtx))
		{
			restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))
			restaurant.GET("/:id", ginrestaurant.FindRestaurant(appCtx))

			//get food if restaurant
			restaurant.GET("/:id/food", ginfood.UserListFoodOfRestaurant(appCtx))

			//Like restaurant
			restaurant.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
			restaurant.DELETE("/:id/unlike", ginrestaurantlike.UserUnLikeRestaurant(appCtx))

			//List user like restaurant
			restaurant.GET("/:id/like", ginrestaurantlike.ListUserLikeRestaurant(appCtx))

			//User rating restaurant
			restaurant.POST("/rating", ginrestaurantrating.CreateRestaurantRating(appCtx))
			restaurant.GET("/:id/rating", ginrestaurantrating.ListRestaurantRating(appCtx))
			restaurant.PUT("/rating/:id_rating", ginrestaurantrating.UpdateRestaurantRating(appCtx))
		}

		v1.GET("/my-rating", middleware.RequireAuth(appCtx), ginrestaurantrating.ListMyRating(appCtx))

		category := v1.Group("/category", middleware.RequireAuth(appCtx))
		{
			category.GET("", gincategory.UserListCategory(appCtx))
		}

		userAddress := v1.Group("address", middleware.RequireAuth(appCtx))
		{
			userAddress.POST("", ginuseraddress.CreateUserAddress(appCtx))
			userAddress.PUT("/:id", ginuseraddress.UpdateUserAddress(appCtx))
			userAddress.DELETE("/:id", ginuseraddress.DeleteUserAddress(appCtx))
			userAddress.GET("", ginuseraddress.ListUserAddress(appCtx))
		}

		cart := v1.Group("/cart", middleware.RequireAuth(appCtx))
		{
			cart.POST("", gincart.CreateItemInCart(appCtx))
			cart.PUT("", gincart.UpdateItemInCart(appCtx))
			cart.DELETE("/:food_id", gincart.DeleteAItemInCart(appCtx))
			cart.DELETE("", gincart.DeleteAllItemInCart(appCtx))
			cart.GET("", gincart.ListItemInCart(appCtx))
		}

		userDeviceToken := v1.Group("/user-device-token", middleware.RequireAuth(appCtx))
		{
			userDeviceToken.GET("/my-device", ginuserdevicetoken.FindUserDeviceToken(appCtx))
			userDeviceToken.POST("", ginuserdevicetoken.CreateUserDeviceToken(appCtx))
		}

		order := v1.Group("/order", middleware.RequireAuth(appCtx))
		{
			order.POST("/momo", ginorder.CreateOrderMomo(appCtx))
			order.POST("/crypto", ginorder.CreateOrderCrypto(appCtx))
			order.GET("/preview", ginorder.PreviewOrder(appCtx))
			order.GET("/:order_id", ginorder.FindOrder(appCtx))
			order.GET("", ginorder.ListOrder(appCtx))
			order.GET("/current", ginorder.ListMyCurrentOrder(appCtx))

			//user confirm received
			order.PUT("/:order_id", ginorder.UserConfirmReceived(appCtx))
		}

		v1.GET("/my-profile", middleware.RequireAuth(appCtx), ginuser.GetUserProfile(appCtx))
		//========================================================================================================

		admin := v1.Group("/admin", middleware.RequireAuthAdmin(appCtx))
		{
			adminOwnerRst := admin.Group("/owner-rst")
			{
				adminOwnerRst.GET("", ginrestaurantowner.ListOwnerRestaurant(appCtx))
				adminOwnerRst.GET("/:id", ginrestaurantowner.FindOwnerRstByAdmin(appCtx))
				adminOwnerRst.POST("", ginrestaurantowner.OwnerRestaurantRegister(appCtx))
				adminOwnerRst.PUT("/:id", ginrestaurantowner.AdminUpdateOwnerRst(appCtx))
			}

			adminRestaurant := admin.Group("/restaurant")
			{
				adminRestaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
				adminRestaurant.GET("", ginrestaurant.ListRestaurantForAdmin(appCtx))
				adminRestaurant.GET("/:id", ginrestaurant.FindRestaurantWithoutStatus(appCtx))
				adminRestaurant.PUT("/:id", ginrestaurant.UpdateRestaurantStatus(appCtx))

				adminRestaurant.GET("/:id/food", ginfood.ListFoodOfRestaurant(appCtx))
			}

			adminCategory := admin.Group("/category")
			{
				adminCategory.POST("", gincategory.CreateCategory(appCtx))
				adminCategory.PUT("/:id", gincategory.UpdateCategory(appCtx))
				adminCategory.DELETE("/:id", gincategory.DeleteCategory(appCtx))
				adminCategory.GET("", gincategory.AdminListCategory(appCtx))
			}

			adminUserDeviceToken := admin.Group("/user-device-token")
			{
				adminUserDeviceToken.GET("", ginuserdevicetoken.ListUserDeviceToken(appCtx))
			}

			adminUser := admin.Group("/user")
			{
				adminUser.PUT("/:id", ginuser.AdminUpdateUser(appCtx))
				adminUser.GET("/:id", ginuser.AdminFindUser(appCtx))
			}

			stats := admin.Group("/stats")
			{
				stats.GET("/overview", ginstatistic.GetOverview(appCtx))
				stats.GET("/order", ginstatistic.GetStatsOrder(appCtx))
				stats.GET("/user", ginstatistic.GetStatsUser(appCtx))
				stats.GET("/food", ginstatistic.GetStatsFood(appCtx))
				stats.GET("/top-user-by-order", ginstatistic.GetStatsTopUserByOrder(appCtx))
			}
		}

		ownerRestaurant := v1.Group("/owner-restaurant")
		{
			ownerRestaurant.POST("/login", ginrestaurantowner.OwnerRestaurantLogin(appCtx))
			ownerRestaurant.POST("/active", ginrestaurantowner.OwnerRestaurantActive(appCtx))
			ownerRestaurant.POST("/send-otp", ginrestaurantowner.SendOTPActiveOwnerRestaurant(appCtx))

			ownerRestaurant1 := ownerRestaurant.Group("/restaurant", middleware.RequireAuthOwnerRestaurant(appCtx))
			{
				ownerRestaurant1.POST("", ginrestaurant.CreateRestaurant(appCtx))
				ownerRestaurant1.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
				ownerRestaurant1.GET("", ginrestaurant.ListRestaurantOwner(appCtx))
				ownerRestaurant1.GET("/:id/food", ginfood.ListFoodOfRestaurant(appCtx)) // Get food of restaurant

				ownerRestaurant1.GET("/:id/order", ginorder.ListOrderRestaurant(appCtx))
				ownerRestaurant1.GET("/:id/order/current", ginorder.ListCurrentOrderRestaurant(appCtx))
				ownerRestaurant1.GET("/order/:order_id", ginorder.FindOrderOfRestaurant(appCtx))
			}

			restaurantOrder := ownerRestaurant.Group("/order", middleware.RequireAuthOwnerRestaurant(appCtx))
			{
				restaurantOrder.PUT("/:order_id", ginorder.RstConfirmPrepareDone(appCtx))
			}

			ownerFood := ownerRestaurant.Group("/food", middleware.RequireAuthOwnerRestaurant(appCtx))
			{
				ownerFood.POST("", ginfood.CreateFood(appCtx))
				ownerFood.PUT("/:id", ginfood.UpdateRestaurant(appCtx))
				ownerFood.DELETE("/:id", ginfood.DeleteRestaurant(appCtx))
				ownerFood.GET("", ginfood.ListFoodOfRestaurant(appCtx))
			}

			ownerCategory := ownerRestaurant.Group("/category")
			{
				ownerCategory.GET("", gincategory.UserListCategory(appCtx))
			}
		}

	}
}
