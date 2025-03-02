package main

import (
	"github.com/dean0731/tencent-cli-go/app"
	"github.com/dean0731/tencent-cli-go/config"
	"github.com/dean0731/tencent-cli-go/log"
)

func main() {
	rootCmd := app.GetRootCmd()
	config.BindFlags(rootCmd)
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		//panic(exception.UnknownError)
	}
}
