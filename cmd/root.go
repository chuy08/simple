/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"simple/internal/logger"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"github.com/spf13/viper"
)

var (
	cfgFile   string = "simple"
	envPrefix string = "simple"
	log       *zap.Logger

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "simple",
		Short: "A Simple Echo Server",
		Long:  `A Golang Echo Server to test multiple gateways.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log = logger.InitLogger()
			return initConfig(cmd)
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("cmd: %v, args: %v", cmd, args)
			log.Sugar().Info("ready set go...")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./simple.json)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command) error {
	v := viper.New()
	v.SetConfigName(cfgFile)
	v.AddConfigPath(".")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable STING_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix(envPrefix)

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	v.AutomaticEnv() // read in environment variables that match

	bindFlags(cmd, v)
	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		//if replaceHyphenWithCamelCase {
		//	configName = strings.ReplaceAll(f.Name, "-", "")
		//}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		//   debug            info
		if f.Changed && v.IsSet(configName) {
			val, _ := cmd.Flags().GetString(configName)
			v.Set(configName, val)
		}
	})
}
