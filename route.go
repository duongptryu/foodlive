package main

import (
	"fooddelivery/component"
	"fooddelivery/middleware"
	"fooddelivery/modules/authsso/fbsso/fbssotransport/ginfbsso"
	"fooddelivery/modules/authsso/googlesso/googlessotransport/gingooglesso"
	"fooddelivery/modules/category/categorytransport/gincategory"
	"fooddelivery/modules/food/foodtransport/ginfood"
	"fooddelivery/modules/restaurant/restauranttransport/ginrestaurant"
	"fooddelivery/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"fooddelivery/modules/restaurantowner/restaurantownertransport/ginrestaurantowner"
	"fooddelivery/modules/upload/uploadtransport/ginupload"
	"fooddelivery/modules/user/usertransport/ginuser"
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

		v1.POST("/upload", ginupload.Upload(appCtx))

		sso := v1.Group("/sso")
		{
			sso.POST("/register-google", gingooglesso.UserGoogleLogin(appCtx))
			sso.POST("/login-google", gingooglesso.UserGoogleLogin(appCtx))

			sso.POST("/register-facebook", ginfbsso.UserFacebookRegister(appCtx))
			sso.POST("/login-facebook", ginfbsso.UserFacebookRegister(appCtx))
		}

		admin := v1.Group("/admin")
		{
			admin.POST("/register-owner-restaurant", ginrestaurantowner.OwnerRestaurantRegister(appCtx))

			restaurant := admin.Group("/restaurant")
			{
				restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
				restaurant.PUT("/:id/status", ginrestaurant.UpdateRestaurantStatus(appCtx))
			}

			category := admin.Group("/category")
			{
				category.POST("", gincategory.CreateCategory(appCtx))
				category.PUT("/:id", gincategory.UpdateCategory(appCtx))
				category.DELETE("/:id", gincategory.DeleteCategory(appCtx))
				category.GET("", gincategory.AdminListCategory(appCtx))
			}
		}

		ownerRestaurant := v1.Group("/owner-restaurant")
		{
			ownerRestaurant.POST("/login", ginrestaurantowner.OwnerRestaurantLogin(appCtx))
			ownerRestaurant.POST("/active", ginrestaurantowner.OwnerRestaurantActive(appCtx))
			ownerRestaurant.POST("/send-otp", ginrestaurantowner.SendOTPActiveOwnerRestaurant(appCtx))

			restaurant := ownerRestaurant.Group("/restaurant", middleware.RequireAuthOwnerRestaurant(appCtx))
			{
				restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
				restaurant.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))

			}

			food := ownerRestaurant.Group("/food")
			{
				food.POST("", ginfood.CreateFood(appCtx))
				food.PUT("/:id", ginfood.UpdateRestaurant(appCtx))
				food.DELETE("/:id", ginfood.DeleteRestaurant(appCtx))
				food.GET("", ginfood.ListFoodOfRestaurant(appCtx))
			}
		}

		restaurant := v1.Group("/restaurant", middleware.RequireAuth(appCtx))
		{
			restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))
			restaurant.GET("/:id", ginrestaurant.FindRestaurant(appCtx))

			//get food if restaurant
			restaurant.GET("/:id/food", ginfood.ListFoodOfRestaurant(appCtx))

			//Like restaurant
			restaurant.POST(":id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
			restaurant.DELETE(":id/unlike", ginrestaurantlike.UserUnLikeRestaurant(appCtx))

			//List user like restaurant
			restaurant.GET(":id/like", ginrestaurantlike.ListUserLikeRestaurant(appCtx))
		}

	}
}
