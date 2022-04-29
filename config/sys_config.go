package config

import (
	"github.com/micrease/micrease-core/config"
	"os"
	"sync"
)

//继承CommonConfig此基础上扩展
type SysConfig struct {
	config.CommonConfig `yaml:",inline"`
}

var once sync.Once
var sysConfig *SysConfig

func InitSysConfig() *SysConfig {
	once.Do(func() {
		var err error
		sysConfig, err = config.LoadConfig[SysConfig]()
		if err != nil {
			os.Exit(1)
		}
	})
	return sysConfig
}

func Get() *SysConfig {
	return sysConfig
}
