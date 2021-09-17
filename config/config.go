// 基础环境变量配置
package config

import (
	"fmt"
	"github.com/yyyThree/gin/helper"
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
	Writer   string `mapstructure:"writer"`
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
			loadEnv()
			fmt.Println("配置文件发生变动：", e.Name, Config)
		})

		err = viper.Unmarshal(&Config)
		unmarshalError(err)
		loadEnv()
	})
}

// 是否是测试环境
func IsDev() bool {
	return Config.App.Env == "debug"
}

// 加载环境变量
func loadEnv() {
	helper.SetEnv(&Config.App.Env, "GIN_ENV")
	helper.SetEnvInt(&Config.Http.Port, "GIN_PORT")
	helper.SetEnvInt(&Config.Http.ReadTimeout, "GIN_READ_TIME_OUT")
	helper.SetEnvInt(&Config.Http.WriteTimeout, "GIN_WRITE_TIME_OUT")
	helper.SetEnvInt(&Config.Http.ShutdownTimeOut, "GIN_SHUTDOWN_TIME_OUT")
	helper.SetEnv(&Config.App.TokenSecret, "GIN_TOKEN_SECRET")
	helper.SetEnv(&Config.Log.Writer, "GIN_LOG_OUT")
	helper.SetEnv(&Config.Log.Dir, "GIN_LOG_DIR")
	helper.SetEnv(&Config.Log.RedisKey, "GIN_LOG_REDIS_KEY")

	helper.SetEnv(&Config.Db.Master.Host, "DB_MASTER_HOST")
	helper.SetEnv(&Config.Db.Master.Port, "DB_MASTER_PORT")
	helper.SetEnv(&Config.Db.Master.User, "DB_MASTER_USER")
	helper.SetEnv(&Config.Db.Master.Password, "DB_MASTER_PASSWORD")
	helper.SetEnvInt(&Config.Db.Master.MaxOpenConnections, "DB_MASTER_MAX_OPEN_CONNECTIONS")
	helper.SetEnvInt(&Config.Db.Master.MaxIdleConnections, "DB_MASTER_MAX_IDLE_CONNECTIONS")
	helper.SetEnvInt(&Config.Db.Master.MaxConnectionIdleTime, "DB_MASTER_MAX_CONNECTION_IDLE_TIME")

	helper.SetEnv(&Config.Db.Slave.Host, "DB_SLAVE_HOST")
	helper.SetEnv(&Config.Db.Slave.Port, "DB_SLAVE_PORT")
	helper.SetEnv(&Config.Db.Slave.User, "DB_SLAVE_USER")
	helper.SetEnv(&Config.Db.Slave.Password, "DB_SLAVE_PASSWORD")
	helper.SetEnvInt(&Config.Db.Slave.MaxOpenConnections, "DB_SLAVE_MAX_OPEN_CONNECTIONS")
	helper.SetEnvInt(&Config.Db.Slave.MaxIdleConnections, "DB_SLAVE_MAX_IDLE_CONNECTIONS")
	helper.SetEnvInt(&Config.Db.Slave.MaxConnectionIdleTime, "DB_SLAVE_MAX_CONNECTION_IDLE_TIME")

	helper.SetEnv(&Config.Redis.Address, "REDIS_ADDRESS")
	helper.SetEnvInt(&Config.Redis.DB, "REDIS_DB")
	helper.SetEnvInt(&Config.Redis.ConnectTimeout, "REDIS_CONNECT_TIMEOUT")
	helper.SetEnvInt(&Config.Redis.ReadTimeout, "REDIS_READ_TIMEOUT")
	helper.SetEnvInt(&Config.Redis.WriteTimeout, "REDIS_WRITE_TIMEOUT")
	helper.SetEnvInt(&Config.Redis.PoolSize, "REDIS_POOL_SIZE")

	helper.SetEnv(&Config.Rabbitmq.Host, "RABBITMQ_HOST")
	helper.SetEnvInt(&Config.Rabbitmq.Port, "RABBITMQ_PORT")
	helper.SetEnv(&Config.Rabbitmq.User, "RABBITMQ_USER")
	helper.SetEnv(&Config.Rabbitmq.Password, "RABBITMQ_PASSWORD")
	helper.SetEnv(&Config.Rabbitmq.Vhost, "RABBITMQ_VHOST")
	helper.SetEnv(&Config.Rabbitmq.AdminUser, "RABBITMQ_ADMIN_USER")
	helper.SetEnv(&Config.Rabbitmq.AdminPassword, "RABBITMQ_ADMIN_PASSWORD")
	helper.SetEnv(&Config.Rabbitmq.ExDirect, "RABBITMQ_EX_DIRECT")
	helper.SetEnv(&Config.Rabbitmq.ExTopic, "RABBITMQ_EX_TOPIC")
	helper.SetEnv(&Config.Rabbitmq.ExDeathLetter, "RABBITMQ_EX_DEATH_LETTER")
	helper.SetEnvInt(&Config.Rabbitmq.TtlQueueMsg, "RABBITMQ_TTL_QUEUE_MSG")
	helper.SetEnvInt(&Config.Rabbitmq.TtlMsg, "RABBITMQ_TTL_MSG")
	helper.SetEnv(&Config.Rabbitmq.LogDir, "RABBITMQ_LOG_DIR")

	fmt.Println("LoadEnv", Config)
}
