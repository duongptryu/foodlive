package main

import (
	"fooddelivery/component"
	"fooddelivery/middleware"
	"fooddelivery/modules/authsso/fbsso/fbssotransport/ginfbsso"
	"fooddelivery/modules/authsso/googlesso/googlessotransport/gingooglesso"
	"fooddelivery/modules/restaurant/restauranttransport/ginrestaurant"
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
			admin.POST("/owner-restaurant/register", ginrestaurantowner.OwnerRestaurantRegister(appCtx))
		}

		ownerRestaurant := v1.Group("/owner-restaurant")
		{
			ownerRestaurant.POST("/login", ginrestaurantowner.OwnerRestaurantLogin(appCtx))
			ownerRestaurant.POST("/active", ginrestaurantowner.OwnerRestaurantActive(appCtx))
			ownerRestaurant.POST("/send-otp", ginrestaurantowner.SendOTPActiveOwnerRestaurant(appCtx))

			restaurant := ownerRestaurant.Group("/restaurant", middleware.RequireAuthOwnerRestaurant(appCtx))
			{
				restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))
			}
		}

		restaurant := v1.Group("/restaurant", middleware.RequireAuth(appCtx))
		{
			restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))
			restaurant.GET("/:id", ginrestaurant.FindRestaurant(appCtx))
		}
	}
}
