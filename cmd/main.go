package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/qibobo/webgo/config"
	"github.com/qibobo/webgo/db/sqldb"
	"github.com/qibobo/webgo/logging"
	"github.com/qibobo/webgo/server"
	"go.uber.org/zap"
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

	logger, err := logging.NewLogger("webgo")
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed to init logger : %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("=======logger: %v\n", logger)
	demodb, err := sqldb.NewDemoSQLDB(conf.DB.DemoDB, *logger.Named("demodb"))
	if err != nil {
		logger.Error("failed to connect to demodb", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("creating http server", zap.Int("port", conf.Server.Port))
	app := server.CreateServer(logger, demodb)

	app.Listen(fmt.Sprintf(":%d", conf.Server.Port))
}
