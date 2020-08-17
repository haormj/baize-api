module github.com/haormj/baize-api

go 1.14

require (
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190808125512-07798873deee
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190307165228-86c17b95fcd5
	github.com/golang/protobuf v1.4.2
	github.com/haormj/ccutil/module/goapi/vision v0.0.0-20200519235848-f4123c31c87d
	github.com/haormj/log v1.0.0
	github.com/haormj/util v1.0.0
	github.com/haormj/version v1.0.0
	github.com/micro/go-micro/v2 v2.7.1-0.20200527112433-192f548c8304
	github.com/spf13/viper v1.7.0
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.24.0
	gopkg.in/yaml.v2 v2.2.4
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
