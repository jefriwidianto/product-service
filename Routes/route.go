package Routes

import (
	"fmt"
	"github.com/labstack/echo"
	"product-service/Controller"
	logger "product-service/Logger"
)

type Routes struct {
	Controller Controller.Controller
	Log        *logger.Logger
}

func (app *Routes) CollectRoutes(e *echo.Echo, host, port string) {
	appRoutes := e

	appRoutes.GET("/list-product", app.Controller.ListProduct)
	appRoutes.GET("/detail-product/:id", app.Controller.DetailProduct)

	app.Log.Fatal(appRoutes.Start(fmt.Sprintf("%s:%s", host, port)).Error())
}
