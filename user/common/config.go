package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v4"
	"go-micro.dev/v4/config"
	"strconv"
)

// GetConsulConfig 注册中心配置
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	source := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	newConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	err = newConfig.Load(source)
	return newConfig, err
}
