package main

import (
	"github.com/dean0731/dean-tool/log"
	"github.com/dean0731/tencent-cli-go/app"
	"github.com/dean0731/tencent-cli-go/config"
)

func main() {
	rootCmd := app.GetRootCmd()
	config.BindFlags(rootCmd)
	rootCmd.AddCommand(VERSIONCMD)
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		//panic(exception.UnknownError)
	}
}
