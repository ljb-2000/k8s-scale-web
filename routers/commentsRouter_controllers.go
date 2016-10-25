package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"] = append(beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"],
		beego.ControllerComments{
			Method: "ReplicationController",
			Router: `/rc`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"] = append(beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"],
		beego.ControllerComments{
			Method: "AutoScale",
			Router: `/as`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"] = append(beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"],
		beego.ControllerComments{
			Method: "RemoveScale",
			Router: `/as/remove/:namespace/:name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"] = append(beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"],
		beego.ControllerComments{
			Method: "NewScale",
			Router: `/as/new`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"] = append(beego.GlobalControllerRouter["github.com/App1905/k8s-scale-web/controllers:APIControllerController"],
		beego.ControllerComments{
			Method: "NewTickScale",
			Router: `/as/newtick`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
