// 基础环境变量配置
package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type configList struct {
	App      app       `mapstructure:"app"`
	Http     http      `mapstructure:"http"`
	Db       databases `mapstructure:"db"`
	Redis    redis     `mapstructure:"redis"`
	Log      log       `mapstructure:"log"`
	Rabbitmq rabbitmq  `mapstructure:"rabbitmq"`
}

type app struct {
	Env         string `mapstructure:"env"`
	TokenSecret string `mapstructure:"token_secret"`
}

type http struct {
	Port            int `mapstructure:"port"`
	ReadTimeout     int `mapstructure:"read_time_out"`
	WriteTimeout    int `mapstructure:"write_time_out"`
	ShutdownTimeOut int `mapstructure:"shutdown_time_out"`
}

type Database struct {
	Host                  string `mapstructure:"host"`
	Port                  string `mapstructure:"port"`
	User                  string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	MaxOpenConnections    int    `mapstructure:"max_open_connections"`
	MaxIdleConnections    int    `mapstructure:"max_idle_connections"`
	MaxConnectionIdleTime int    `mapstructure:"max_connection_idle_time"`
}

type databases struct {
	Master Database `mapstructure:"master"`
	Slave  Database `mapstructure:"slave"`
}

type redis struct {
	Address        string `mapstructure:"address"`
	Password       string `mapstructure:"password"`
	DB             int    `mapstructure:"db"`
	ConnectTimeout int    `mapstructure:"connect_timeout"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	PoolSize       int    `mapstructure:"pool_size"`
}

type log struct {
	Out      string `mapstructure:"out"`
	Dir      string `mapstructure:"dir"`
	RedisKey string `mapstructure:"redisKey"`
}

type rabbitmq struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	Vhost         string `mapstructure:"vhost"`
	AdminUser     string `mapstructure:"admin_user"`
	AdminPassword string `mapstructure:"admin_password"`
	ExDirect      string `mapstructure:"ex_direct"`
	ExTopic       string `mapstructure:"ex_topic"`
	ExDeathLetter string `mapstructure:"ex_death_letter"`
	TtlQueueMsg   int    `mapstructure:"ttl_queue_msg"`
	TtlMsg        int    `mapstructure:"ttl_msg"`
	LogDir        string `mapstructure:"log_dir"`
}

var (
	load   sync.Once // 确保配置文件仅需加载一次
	Config configList
)

// 加载配置文件，仅需加载一次
func Load() {
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

// 是否是测试环境
func IsDev() bool {
	return Config.App.Env == "debug"
}
