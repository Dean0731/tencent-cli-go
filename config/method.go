package config

import (
	"github.com/dean0731/tencent-cli-go/exception"
	"github.com/dean0731/tencent-cli-go/log"
	"github.com/dean0731/tencent-cli-go/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func SetKeyValue(key string, value any) {
	viper.Set(key, value)
}

func SetKeyValueIfBlank(key string, value any) {
	if GetString(key) == "" {
		SetKeyValue(key, value)
	}
	printConfig(key)
}

func GetStringWithDefault(key string, defaultValue string) string {
	value := GetString(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func printConfig(k string) {
	if k == "" {
		return
	}
	if value := GetString(k); value != "" {
		log.Debugf("%s: %v", k, value)
	}
}

func BindFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool(OutPutJsonFormat, false, "输出JSON格式化")
	cmd.PersistentFlags().Bool(Debug, false, "日志阶级，是否打印参数")
	for _, key := range GlobalStringParam {
		cmd.PersistentFlags().String(key, utils.GetEnv(key), key)
	}

	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		panic(exception.UnknownError.SetMessage(err.Error()))
	}
	LoadConfig()
	viper.SetDefault(Region, DefaultRegion)
}

func PrintConfig() {
	printConfig(Region)
	printConfig(EndPoint)
	printConfig(SecretKey)
	printConfig(SecretId)
	printConfig(HttpProxy)
}
