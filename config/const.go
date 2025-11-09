package config

import _const "github.com/dean0731/dean-tool/const"

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
	HttpProxy = _const.HttpProxy
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
	AppDir      = ".tencent-cli"
	GithubOwner = "Dean0731"
	GithubRepo  = "tencent-cli-go"
	GitAppName  = "tencent_cli_go"
)
