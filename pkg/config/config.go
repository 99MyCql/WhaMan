package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Conf 配置数据实体
type Conf struct {
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

// New 读取配置文件，获取配置数据
func New(filepath string) (*Conf, error) {
	// 读取并解析 conf.yaml 文件
	inFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "读取配置文件失败")
	}
	conf := new(Conf)
	err = yaml.Unmarshal(inFile, conf)
	if err != nil {
		return nil, errors.Wrap(err, "解析配置文件失败")
	}
	return conf, nil
}
