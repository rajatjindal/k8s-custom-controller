package client

import (
	"github.com/rajatjindal/k8s-custom-controller/config"
	"github.com/rajatjindal/k8s-custom-controller/pkg/controller"
)

func Run(conf *config.Config) {
	eventHandler := conf.Discovery
	controller.Start(conf, eventHandler)
}
