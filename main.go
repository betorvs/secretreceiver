package main

import (
	"github.com/betorvs/secretreceiver/config"
	"github.com/betorvs/secretreceiver/controller"
	_ "github.com/betorvs/secretreceiver/gateway/kubeclient"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := echo.New()
	g := e.Group("/secretreceiver/v1")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/health", controller.CheckHealth)
	g.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	g.GET("/secret/:namespace/:name", controller.CheckSecret)
	g.POST("/secret", controller.CreateSecret)
	g.PUT("/secret", controller.UpdateSecret)
	g.DELETE("/secret/:namespace/:name", controller.DeleteSecret)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
