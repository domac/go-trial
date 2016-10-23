package proxy

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ProxyConfig struct {
	Bind         string   `yaml:"bind"`
	WaitQueueLen int      `yaml:"wait_queue_len"`
	MaxConn      int      `yaml:"max_conn"`
	Timeout      int      `yaml:"timeout"`
	FailOver     int      `yaml:"failover"`
	Backend      []string `yaml:"backend"`
	Stats        string   `yaml:"stats"`
}

//解析配置文件
func ParseConfigFile(filepath string) (*ProxyConfig, error) {
	pconfig := ProxyConfig{}
	if config, err := ioutil.ReadFile(filepath); err == nil {

		if err = yaml.Unmarshal(config, &pconfig); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	return &pconfig, nil
}
