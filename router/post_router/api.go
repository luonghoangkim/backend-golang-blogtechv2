package postrouter

import (
	"backend-blogtechv2/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	PostHandler handler.PostHandler
}

func (api *API) SetupRouter() {
	//user
	api.Echo.POST("/admin/post", api.PostHandler.HandlePost) 
}