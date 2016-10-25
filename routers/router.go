package routers

import (
	"github.com/App1905/k8s-scale-web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api", beego.NSInclude(&controllers.APIControllerController{}))
	beego.AddNamespace(ns)
}
