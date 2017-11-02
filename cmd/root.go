// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/rajatjindal/k8s-custom-controller/pkg/client"
	"github.com/rajatjindal/k8s-custom-controller/config"
	"github.com/Sirupsen/logrus"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "k8s-custom-controller",
	Short: "A k8s controller that react on anything with annotation test.alpha.kubernetes.io/service-name",
	Long: `A k8s controller which react as soon as a new service is created with annotation
and takes some action on it.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Validate()
		if err != nil {
			logrus.Fatal(err)
		}

		client.Run(conf)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.k8s-custom-controller.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(config.ConfigFileName)
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}