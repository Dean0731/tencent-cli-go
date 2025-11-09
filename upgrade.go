// selfupdate.go
package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/dean0731/dean-tool/github"
	"github.com/dean0731/dean-tool/log"
	"github.com/dean0731/dean-tool/utils"
	"github.com/dean0731/tencent-cli-go/config"
	"github.com/spf13/cobra"
)

var UpgradeCMD = &cobra.Command{
	Use:   "upgrade",
	Short: "更新程序",
	Long:  "下载最新发布版本，更新程序",
	Run: func(cmd *cobra.Command, args []string) {
		selfUpdate()
	}}

func selfUpdate() {
	latestTag := github.GetLatestReleaseTag(config.GithubOwner, config.GithubRepo)

	currentVer := strings.TrimPrefix(VERSION, "v")
	latestVer := strings.TrimPrefix(latestTag, "v")

	if currentVer == latestVer {
		log.Println("Already up to date.")
		return
	}

	log.Printf("New version available: %s (current: %s)", latestTag, VERSION)
	// 下载存放路径
	downloadDir := filepath.Join(".", "download")
	// 解压后路径
	folder := config.GitAppName + "-" + latestTag + "-" + runtime.GOOS + "-" + runtime.GOARCH
	filename := folder + config.GithubReleaseExt
	url := "https://github.com/" + config.GithubOwner + "/" + config.GithubRepo + "/releases/latest/download/" + filename
	utils.ExtractExecutable(utils.HttpGetDownloadWithProgress(url, downloadDir, filename), downloadDir)
	err := os.Rename(filepath.Join(downloadDir, folder, config.GithubProcessName), utils.GetCurrentExecutablePath())
	if err != nil {
		panic(err)
	}

	log.Println("Binary updated successfully.")
}
