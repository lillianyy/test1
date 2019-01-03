package models

import (
	"time"
	"github.com/lillion-y/test1/tree/master/istio-ui-master-2/src/pkg"
    "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var controller *pkg.Controller

func InitController() {

	options := pkg.ControllerOptions{
		DomainSuffix : "cluster.local",
		ResyncPeriod : 60*time.Second,
		WatchedNamespace : "",
	}

	controller = pkg.NewController(pkg.GetKubeClent(), options)
}

func Run(stop <-chan struct{})  {
	go controller.Run(stop)
}

func DeploysList(deployIndexs []string, namespace string) []interface{} {
	return controller.GetDeployList(deployIndexs, namespace)
}


func ListKeys() []string {
	return controller.ListKeys()
}

func GetByKey(key string) (item interface{}, exists bool, err error) {
	return controller.GetByKey(key)
}