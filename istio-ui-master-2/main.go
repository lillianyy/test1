package main

import (
	_ "github.com/lillion-y/test1/tree/master/istio-ui-master-2/src/routers"
	"github.com/astaxie/beego"
	"github.com/lillion-y/test1/tree/master/istio-ui-master-2/src/models"
	"github.com/lillion-y/test1/tree/master/istio-ui-master-2/src/pkg"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	stop := make(chan struct{})
	pkg.InitKubeClient()
	pkg.InitConfigClient()
	pkg.InitDeployIndexStore()
	models.InitController()
	models.Run(stop)
	
	beego.Run()
	close(stop)
}
