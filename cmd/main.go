package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/qibobo/webgo/config"
	"github.com/qibobo/webgo/server"
)

func main() {
	var path string
	flag.StringVar(&path, "c", "", "config file")
	flag.Parse()
	if path == "" {
		fmt.Fprintln(os.Stderr, "missing config file")
		os.Exit(1)
	}

	configFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to open config file '%s' : %s\n", path, err.Error())
		os.Exit(1)
	}
	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to read config file '%s' : %s\n", path, err.Error())
		os.Exit(1)
	}
	var conf *config.Config
	conf, err = config.LoadConfig(configBytes)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to read config: '%s'\n", err.Error())
		os.Exit(1)
	}
	configFile.Close()

	err = conf.Validate()
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to validate configuration : %s\n", err.Error())
		os.Exit(1)
	}
	app := server.CreateServer()
	app.Listen(fmt.Sprintf(":%d", conf.Port))
}
