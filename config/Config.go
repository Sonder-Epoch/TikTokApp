package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `config.yaml:"port"`
}

var Conf = new(Config)

func Init() {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("配置文件不存在:%s", err))
		} else {
			panic(fmt.Errorf("配置文件加载错误:%s", err))
		}
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("配置文件解析错误:%s", err))
	}
}
