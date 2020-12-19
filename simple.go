package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

// Variables are set at build using -ldflags
var buildTime = "now"
var buildSha5 = "sha5"
var buildVersion = "development"

var (
	flgVersion bool
)

func main() {
	port := flag.String("port", "1323", "Server Port")
	flag.BoolVar(&flgVersion, "version", false, "application version")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("   simple [command]\n")
		fmt.Printf("\n")
		fmt.Printf("Available Commands:\n")
		fmt.Printf("   version     print the client version information\n")
		fmt.Printf("\n")
		fmt.Printf("Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	fmt.Println(flag.Args())

	if flgVersion {
		fmt.Printf("Version: %s\nBuild on %s from %s\n", buildVersion, buildTime, buildSha5)
		os.Exit(0)
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", *port),
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
	log.Info(msg)
	return c.String(200, msg)
}
