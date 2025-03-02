# 命令行工具简介
欢迎使用腾讯云命令行工具(TCCLI)，TCCLI是管理腾讯云资源的统一工具。通过腾讯云命令行工具，您可以快速轻松的调用腾讯云 API来管理您的腾讯云资源。您还可以基于腾讯云的命令行工具来做自动化和脚本处理，能够以更多样的方式进行组合和重用。
# 安装TCCLI
1. 下载最新版本的即可二进制文件即可，注意选择对应的架构
# 配置TCCLI
1. 在家目录下创建./tencent-cli/Config.yaml
```yaml
SecretId: AKID*****************************
SecretKey: Key*****************************
# HttpProxy: http://127.0.0.1:1080
# Region: ap-guangzhou
# Endpoint: cvm.tencentcloudapi.com
```
2. 执行命令时添加参数
```bash
tencent_linux cvm DescribeInstances --Region ap-guangzhou
```
3. 设置环境变量

注意：如果环境变量定义了相关配置，则会优先于配置文件生效。分别为 参数指定->环境变量->配置文件->默认值

# 使用TCCLI
命令行工具集成了腾讯云所有支持云 API 的产品，可以在命令行下完成对腾讯云产品的配置和管理。包括使用TCCLI创建云服务器，操作云服务器，通过TCCLI创建CBS盘、查看CBS盘使用情况，通过TCCLI创建VPC网络、往VPC网络中添加资源等等，所有在控制台页面能完成的操作，均能再命令行工具上执行命令实现。
* 通过tccli cvm DescribeInstances命令查看当前账号有哪些云服务器。
* 通过tccli cbs DescribeDisks命令查看有CBS盘列表。

以创建一台cvm为例(**请注意demo中非简单类型的参数必须为标准json格式**)：
```bash
tccli cvm RunInstances --InstanceChargeType POSTPAID_BY_HOUR --InstanceChargePrepaid '{"Period":1,"RenewFlag":"DISABLE_NOTIFY_AND_MANUAL_RENEW"}'
 --Placement '{"Zone":"ap-guangzhou-2"}' --InstanceType S1.SMALL1 --ImageId img-8toqc6s3 --SystemDisk '{"DiskType":"CLOUD_BASIC", "DiskSize":50}'
--InternetAccessible '{"InternetChargeType":"TRAFFIC_POSTPAID_BY_HOUR","InternetMaxBandwidthOut":10,"PublicIpAssigned":true}' --InstanceCount 1
--InstanceName TCCLI-TEST --LoginSettings '{"Password":"isd@cloud"}' --SecurityGroupIds '["sg-0rszg2vb"]' --HostName TCCLI-HOST-NAME1
```
更多功能，您可以通过tccli help查看支持的产品，通过tccli cvm help（以cvm举例）查看产品支持的接口。

通过tccli cbs DescribeDisks help(以cbs产品的DescribeDisks接口为例) 查看接口支持的参数。



