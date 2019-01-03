package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lillion-y/test1/tree/master/istio-ui-master-2/src/pkg"
)

type TemplateController struct {
	beego.Controller
}


/**
get istio config from local file
 */
func (this *TemplateController) GetMeshConfig() {

	config, err := pkg.GetMeshConfig()
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code": -1, "msg": err.Error(), "data": nil}
		this.ServeJSON()
	}

	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data" : config}
	this.ServeJSON()
}

/**
write to local file and post to k8s
 */
func (this *TemplateController) GetInjectConfig() {

	config, err := pkg.GetInjectConfigFromConfigMap()
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code": -1, "msg": err.Error(), "data": nil}
		this.ServeJSON()
	}

	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data" : config}
	this.ServeJSON()
}