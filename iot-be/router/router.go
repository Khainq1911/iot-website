package router

import (
	"web-ivsr-be/handlers"
	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo    *echo.Echo
	Handler handlers.SiteHandler
}

func (api *Api) SetUpRouter() {
	api.Echo.GET("/", api.Handler.GetData)
	api.Echo.GET("/:node_id", api.Handler.GetDataId)
	api.Echo.PATCH("/:node_id", api.Handler.PutData)
}
