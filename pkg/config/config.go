package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Conf 配置数据实体
type Conf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	MysqlUrl string `yaml:"mysqlUrl"`
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