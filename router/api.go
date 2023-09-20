package router

import (
	"backend-blogtechv2/handler"
	"backend-blogtechv2/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	PostHandler handler.PostHandler
}

func (api *API) SetupRouter() {
	// user
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	// profile
	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	user.PUT("/profile/update", api.UserHandler.UpdateProfile) 

	// post
	post := api.Echo.Group("/post", middleware.JWTMiddleware())
	post.POST("", api.PostHandler.HandlePost)
	post.GET("/selectedPost", api.PostHandler.GetPostByID)
	post.GET("/getAllPost", api.PostHandler.GetAllPostsByTable)
	post.PUT("/updatePost", api.PostHandler.UpdatePost)


}