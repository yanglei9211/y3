package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"y3/controller"
)

func InitRouters() {
	nsTest := web.NewNamespace("/api/y3",
		web.NSRouter("/sum", &controller.MainController{}),
		web.NSRouter("item_info", &controller.ItemInfoController{}),
		web.NSRouter("item_details", &controller.ItemDetailsController{}))
	logs.Info("setup routers success")
	web.AddNamespace(nsTest)
	web.Router("/", &controller.WelcomeController{})
}
