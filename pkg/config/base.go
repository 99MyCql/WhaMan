package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Conf 全局配置数据
var Conf *BaseConf

// BaseConf 配置数据实体
type BaseConf struct {
	Env           string `yaml:"env"`
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	MysqlUrl      string `yaml:"mysqlUrl"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	SessionSecret string `yaml:"sessionSecret"`
	SslCert       string `yaml:"sslCert"`
	SslKey        string `yaml:"sslKey"`
}

// Init 读取配置文件，获取配置数据
func Init(filepath string) {
	// 读取并解析 conf.yaml 文件
	inFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	Conf = new(BaseConf)
	err = yaml.Unmarshal(inFile, Conf)
	if err != nil {
		panic(err)
	}
}
