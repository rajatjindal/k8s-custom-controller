package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/rajatjindal/k8s-custom-controller/pkg/discovery"
)

var ConfigFileName = ".pfpt-k8s-discovery"
const DefaultAnnotation = "rajatjindal/test/"

type K8sSearcher struct {
	Annotation string `json:"annotation"`
}

type Config struct {
	Discovery discovery.Discovery `json:"discovery"`
	K8sSearcher K8sSearcher `json:"k8sSearcher"`
}

func Validate() (*Config, error) {
	conf := &Config{}

	if viper.GetString("discovery.url") == "" {
		return nil, fmt.Errorf("discovery.url is mandatory")
	}
	conf.Discovery.Url = viper.GetString("discovery.url")

	if viper.GetString("discovery.environment") == "" {
		return nil, fmt.Errorf("discovery.environment is mandatory")
	}
	conf.Discovery.Environment= viper.GetString("discovery.environment")

	if viper.GetString("k8sSearcher.annotation") == "" {
		conf.K8sSearcher.Annotation = DefaultAnnotation
	}

	return conf, nil
}


