package router

import (
	"mk/spider_a/server/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/spiderapi", &controller.SpiderController{}, "post:SpiderApi")
}
