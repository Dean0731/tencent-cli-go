package app

import (
	"embed"
	"encoding/json"
	"github.com/dean0731/dean-tool/exception"
	"github.com/dean0731/tencent-cli-go/app/common"
	"github.com/spf13/cobra"
	"io/fs"
)

//go:embed data/*
var embeddedFiles embed.FS

func GetRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Tencent CLI",
		Short: "腾讯云CLI",
		Long:  `TCCLI是管理腾讯云资源的统一工具。通过腾讯云命令行工具，您可以快速轻松的调用腾讯云 API来管理您的腾讯云资源。您还可以基于腾讯云的命令行工具来做自动化和脚本处理，能够以更多样的方式进行组合和重用。`,
	}
	fs.WalkDir(embeddedFiles, ".", func(path string, d fs.DirEntry, _ error) error {
		if !d.IsDir() {
			content, _ := fs.ReadFile(embeddedFiles, path)
			var api common.Api
			err := json.Unmarshal(content, &api)
			if err != nil {
				panic(exception.JsonLoadError.SetMessage(err.Error()))
			}
			rootCmd.AddCommand(common.LoadApi(&api))
		}
		return nil
	})
	return rootCmd
}
