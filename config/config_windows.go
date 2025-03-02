//go:build windows

package config

import (
	"errors"
	"github.com/dean0731/tencent-cli-go/exception"
	"github.com/dean0731/tencent-cli-go/utils"
	"github.com/spf13/viper"
	"path/filepath"
)

func LoadConfig() {
	viper.AddConfigPath(".")
	if home := utils.GetEnv(WindowsHomeEnv); home != "" {
		viper.AddConfigPath(filepath.Join(home, AppDir))
	}
	viper.SetConfigName(ConfigurationFileName)
	viper.SetConfigType(ConfigurationFileType)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			panic(exception.FileFormatError.SetMessage(err.Error()))
		}
	}
}
