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
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the Server at a specified port",
	Long:  `Turn the thing on and listen somewhere`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		//fmt.Println("server called")
		//fmt.Println(port)
		start(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().StringP("port", "p", "80", "Port to run server on")
}

func start(port string) {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e := echo.New()
	e.Logger.SetOutput(os.Stdout)
	e.Debug = true
	e.HideBanner = true

	e.GET("/_command/status", getStatus)
	e.GET("/_command/*", getCommand)

	e.Logger.Fatal(e.StartServer(s))
}

func getCommand(c echo.Context) error {
	msg := "Command Application"
	log.Info(msg)
	return c.String(200, msg)
}

func getStatus(c echo.Context) error {
	msg := "Ok"
	log.Debug(msg)
	return c.String(200, msg)
}
