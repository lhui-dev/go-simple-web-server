package configuration

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

// 项目总配置
type config struct {
	Server serverConfig `yaml:"server"`
	Db     dbConfig     `yaml:"db"`
	Log    logConfig    `yaml:"log"`
}

// 项目端口号相关配置
type serverConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

// db相关配置信息
type dbConfig struct {
	RedisConfig   redisConfig   `yaml:"redis"`
	MysqlConfig   mysqlConfig   `yaml:"mysql"`
	MongoDbConfig mongodbConfig `yaml:"mongodb"`
}

// redis相关配置
type redisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// mysql相关配置
type mysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	MaxIdle  uint   `yaml:"maxIdle"`
	MaxOpen  uint   `yaml:"maxOpen"`
}

// mongodb相关配置
type mongodbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// 日志相关配置
type logConfig struct {
	Path     string `yaml:"path"`
	FileName string `yaml:"fileName"`
	Mode     string `yaml:"mode"`
}

// Config 配置对象
var Config *config

func init() {
	// 从项目目录下面开始作为根路径
	var filePath = "./configuration.yaml"

	// 获取当前工作目录
	dir, err := os.Getwd()

	if err != nil {
		fmt.Printf("fail to get current directory: %v\n", err)
	}

	// 拼接完整的文件名
	filename := path.Join(dir, filePath)
	// 打开并读取yaml配置文件
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("open configuration.yaml failed!")
		fmt.Println(err.Error())
		panic(err)
	}

	// 使用yaml中Unmarshal方法，解析yaml配置文件并绑定定义的结构体
	if err = yaml.Unmarshal(yamlFile, &Config); err != nil {
		fmt.Println("yaml configuration file Unmarshal failed!")
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("Load configuration.yaml successfully...")
}
