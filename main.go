package main

import (
	_ "github.com/App1905/k8s-scale-web/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateRight = ">>>"
	beego.SetStaticPath("/assets", "static")
	beego.Run()
}
