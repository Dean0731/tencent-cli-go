package common

import (
	"github.com/dean0731/tencent-cli-go/config"
	"github.com/dean0731/tencent-cli-go/log"
	"github.com/dean0731/tencent-cli-go/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"strings"
)

func GetClient(production string) *common.Client {
	credential := common.NewCredential(config.GetString(config.SecretId), config.GetString(config.SecretKey))
	cpf := profile.NewClientProfile()
	config.SetKeyValueIfBlank(config.EndPoint, production+config.EndPointSub)
	cpf.HttpProfile.Endpoint = config.GetString(config.EndPoint)
	if proxy := config.GetString(config.HttpProxy); proxy != "" {
		cpf.HttpProfile.Proxy = config.GetString(config.HttpProxy)
	}
	return common.NewCommonClient(credential, config.GetString(config.Region), cpf)
}

func getJsonArgs(cmd *cobra.Command) string {
	result := map[string]interface{}{}
	flags := cmd.Flags()
	flags.VisitAll(func(flag *pflag.Flag) {
		name := flag.Name
		if flag.Changed && !utils.ContainsInSlice(config.GlobalStringParam, name) && !utils.ContainsInSlice(config.GlobalBoolParam, name) {
			log.Debugf("Origin Params %s: %s", flag.Name, flag.Value.String())
			if flag.Value.Type() == "stringArray" {
				data := flag.Value.String()
				data = strings.Replace(data, "\"", "", -1)
				data = strings.Replace(data, "'", "\"", -1)
				result[flag.Name] = utils.JsonLoad([]byte(data))
			} else {
				result[flag.Name] = flag.Value
			}
		}
	})
	str := utils.JsonDump(result)
	log.Debugf("Request Params: %s", str)
	return str
}

func LoadApi(api *Api) *cobra.Command {

	var cmd = &cobra.Command{
		Use:     api.MetaData.ServiceShortName,
		Short:   api.MetaData.ServiceNameCN,
		Long:    api.MetaData.ApiBrief,
		Version: api.MetaData.ApiVersion,
	}
	for key, value := range api.Actions {
		subCmd := &cobra.Command{
			Use:   key,
			Short: value.Name,
			Long:  value.Document,
			PersistentPreRun: func(cmd *cobra.Command, args []string) {
				config.PrintConfig()
			},
			Run: func(cmd *cobra.Command, args []string) {
				request := tchttp.NewCommonRequest(api.MetaData.ServiceShortName, api.MetaData.ApiVersion, key)
				err := request.SetActionParameters(getJsonArgs(cmd))
				if err != nil {
					panic(err)
				}
				response := tchttp.NewCommonResponse()
				err = GetClient(api.MetaData.ServiceShortName).Send(request, response)
				if err != nil {
					panic(err)
				}
				log.PrintJson(string(response.GetBody()))
			},
		}
		for _, member := range api.Objects[value.Input].Members {
			if member.Type == "bool" {
				subCmd.Flags().Bool(member.Name, false, member.Document)
			} else if member.Type == "int" {
				subCmd.Flags().Int(member.Name, 0, member.Document)
			} else if member.Type == "list" {
				subCmd.Flags().StringArray(member.Name, nil, member.Document)
			} else {
				subCmd.Flags().String(member.Name, "", member.Document)
			}
			if member.Required {
				subCmd.MarkFlagRequired(member.Name)
			}
		}
		cmd.AddCommand(subCmd)
	}
	return cmd
}
