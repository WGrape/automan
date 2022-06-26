package main

import (
	"automan/api/code"
	"automan/api/server"
	"automan/internal/config"
	"automan/internal/logger"
	"automan/internal/model"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var GlobalConfig *model.Config
var GlobalLogger *zap.SugaredLogger

func init() {
	var err error
	GlobalLogger = logger.New()

	server.Register()

	GlobalConfig, err = config.New()
	if err != nil {
		GlobalLogger.Fatalf("New config error: %v", err)
		os.Exit(code.ErrExit)
		return
	}
}

func main() {
	GlobalLogger.Infof("Automan run success: %#v, %#v", GlobalConfig, GlobalLogger)
	err := http.ListenAndServe(fmt.Sprintf(":%d", GlobalConfig.Port), nil)
	if err != nil {
		GlobalLogger.Fatalf("ListenAndServe error: %v", err)
		os.Exit(code.ErrExit)
		return
	}
}
