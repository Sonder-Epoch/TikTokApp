package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Mysql  MysqlConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	Log    LogConfig    `yaml:"log"`
	OSS    OSSConfig    `yaml:"oss"`
}
type ServerConfig struct {
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}
type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Pwd      string `yaml:"pwd"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	LogLevel int    `yaml:"logLevel"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}
type RedisConfig struct {
	Address  string `yaml:"address"`
	DB       int    `yaml:"DB"`
	PoolSize int    `yaml:"poolSize"`
}
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxSize"`
	MaxAge     int    `json:"maxAge"`
	MaxBackups int    `json:"maxBackups"`
}
type OSSConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
}

var Conf = initConfig()

func initConfig() *Config {
	conf := new(Config)
	//workDir, _ := os.Getwd()
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(workDir + "/config")
	viper.SetConfigFile("E:\\code\\go\\TikTokApp\\config\\config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("配置文件不存在:%s", err))
		} else {
			panic(fmt.Errorf("配置文件加载错误:%s", err))
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		panic(fmt.Errorf("配置文件解析错误:%s", err))
	}
	return conf
}
