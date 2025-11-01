package main

import "github.com/spf13/cobra"
import "github.com/dean0731/dean-tool/log"

var VERSION = "v1.0.0"
var VERSIONCMD = &cobra.Command{
	Use:   "version",
	Short: "版本号",
	Long:  "版本号",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf(VERSION)
	}}
