// 基础环境变量配置
package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type configList struct {
	App      app    `mapstructure:"app"`
	Http     http   `mapstructure:"http"`
	Language string `mapstructure:"language"`
}

type app struct {
	TokenSecret string `mapstructure:"token_secret"`
}

type http struct {
	Port int `mapstructure:"port"`
	ReadTimeout int `mapstructure:"read_time_out"`
	WriteTimeout int `mapstructure:"write_time_out"`
}

var (
	load   sync.Once // 确保配置文件仅需加载一次
	Config configList
)

// 加载配置文件，仅需加载一次
func LoadConfig() {
	load.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("配置文件读取错误: %s \n", err))
		}
		unmarshalError := func(err error) {
			if err == nil {
				return
			}
			panic(fmt.Errorf("配置文件映射错误: %s \n", err))
		}
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			// 重新读取配置文件
			err = viper.Unmarshal(&Config)
			unmarshalError(err)
			fmt.Println("配置文件发生变动：", e.Name, Config)
		})

		err = viper.Unmarshal(&Config)
		unmarshalError(err)
	})
}
