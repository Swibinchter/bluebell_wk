package setting

// 使用viper进行配置文件的读取，存到对应结构体中

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 此处需要在反引号内使用mapstructure的标签才能被viper识别，不管是yaml还是json等
type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// Config 全局变量，结构体实例，保存本项目所有配置信息
var Config = new(AppConfig)

func Init() (err error) {
	// 指定配置文件，包括名字和类型
	viper.SetConfigFile("./conf/config.yaml")

	// 读取配置文件的信息
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}

	// 将读取到的配置信息反序列化存放到结构体实例中
	err = viper.Unmarshal(Config)
	if err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	// 监听配置文件，实现热加载而无需重启
	viper.WatchConfig()
	// 配置文件改动后执行回调函数/钩子函数，重新将配置信息更新到结构体实例中，实现热加载
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已被修改...")
		// 更新配置信息到结构体实例中
		err = viper.Unmarshal(Config)
		if err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})

	return
}
