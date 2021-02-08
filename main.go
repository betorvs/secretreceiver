package main

import (
	"github.com/betorvs/secretreceiver/config"
	"github.com/betorvs/secretreceiver/controller"
	_ "github.com/betorvs/secretreceiver/gateway/customlog"
	_ "github.com/betorvs/secretreceiver/gateway/kubeclient"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
