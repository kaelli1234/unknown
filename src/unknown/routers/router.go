package routers

import (
	"github.com/astaxie/beego"

	"unknown/controllers"
)

func init() {
	beego.Router("/shops", &controllers.ExampleController{}, "get:ShopList")
	beego.Router("/shops", &controllers.ExampleController{}, "post:ShopAdd")
	beego.Router("/votes", &controllers.ExampleController{}, "post:VoteAdd")
	beego.Router("/votes/:id", &controllers.ExampleController{}, "get:VoteGet")
	beego.Router("/votes/post", &controllers.ExampleController{}, "post:VotePost")
	beego.Router("/votes/:id/result", &controllers.ExampleController{}, "get:VoteResult")
}
