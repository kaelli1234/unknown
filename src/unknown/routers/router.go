package routers

import (
	"github.com/astaxie/beego"

	"unknown/controllers"
	"unknown/controllers/v1"
	// "unknown/middlewares"
)

func init() {
	beego.Router("/ping", &controllers.ExampleController{}, "get:Home")
	beego.AddNamespace(beego.NewNamespace("/v1",
		beego.NSRouter("/home", &v1.ExampleController{}, "get:Home"),
		beego.NSRouter("/test", &v1.ExampleController{}, "post:Test"),
		beego.NSRouter("/token", &v1.ExampleController{}, "post:Token"),
		beego.NSRouter("/mysql", &v1.ExampleController{}, "post:Mysql"),
	))
}
