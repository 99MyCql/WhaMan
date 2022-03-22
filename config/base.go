package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Conf 全局基础配置数据
var Conf *base

type mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type ssl struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

// base 基础配置数据实体
type base struct {
	ProjectName   string `yaml:"project_name"`
	Env           string `yaml:"env"`
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	Mysql         *mysql `yaml:"mysql"`
	LogLevel      string `yaml:"log_level"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	SessionSecret string `yaml:"session_secret"`
	Ssl           *ssl   `yaml:"ssl"`
}

// Init 读取配置文件，获取配置数据
func Init(filepath string) {
	// 读取并解析 conf.yaml 文件
	inFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	Conf = new(base)
	err = yaml.Unmarshal(inFile, Conf)
	if err != nil {
		panic(err)
	}
}
