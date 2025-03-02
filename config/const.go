package config

const (
	Debug            = "D"
	OutPutJsonFormat = "J"
)

const (
	// 全局参数key
	Region    = "Region"
	EndPoint  = "EndPoint"
	SecretId  = "SecretId"
	SecretKey = "SecretKey"
	HttpProxy = "HttpProxy"
)

var (
	GlobalStringParam = []string{Region, EndPoint, SecretId, SecretKey, HttpProxy}
)
var (
	GlobalBoolParam = []string{OutPutJsonFormat, Debug}
)

const (
	EndPointSub   = ".tencentcloudapi.com"
	DefaultRegion = "ap-guangzhou"
)

const (
	ConfigurationFileName = "Config"
	ConfigurationFileType = "yaml"

	WindowsHomeEnv = "USERPROFILE"
	LinuxHomeEnv   = "HOME"

	AppDir = ".tencent-cli"
)
