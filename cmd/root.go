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
	"simple/config"
	"simple/internal/logger"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configFile string
	log        *zap.Logger
	serverPort int32
	vi         *viper.Viper

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
			//log.Sugar().Info("ready set go...")
			cfg := config.New(vi)
			log.Sugar().Infof("config foo: %s, port: %d", cfg.GetFoo(), cfg.GetServerPort())
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

	rootCmd.Flags().Int32Var(&serverPort, "port", 80, "Port to run server on")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.json)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command) error {
	vi = viper.New()
	vi.SetConfigName(configFile)
	vi.AddConfigPath(".")

	// If a config file is found, read it in.
	if err := vi.ReadInConfig(); err == nil {
		log.Sugar().Infof("Found config file: %s", vi.ConfigFileUsed())
	}

	vi.AutomaticEnv() // read in environment variables that match

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable STING_NUMBER. This helps
	// avoid conflicts.
	vi.SetEnvPrefix("simple")

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
	vi.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	vi.BindEnv("server.port", "SERVER_PORT")

	vi.BindPFlag("server.port", cmd.Flags().Lookup("port"))

	bindFlags(cmd, vi)

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
