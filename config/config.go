package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configure	struct {
	Mysql	MysqlConfig		`yaml:"mysql"`
	Produce	Produce			`yaml:"produce"`
}

type Produce struct {
	Port 	string	`yaml:"port"`
}


type MysqlConfig struct {
	SqlHost		string	`yaml:"host"`
	SqlPort 	string	`yaml:"port"`
	SqlDatabase	string	`yaml:"database"`
	SqlUsername	string	`yaml:"username"`
	SqlPassword	string	`yaml:"password"`
}

var	Cfg = &Configure{}

func init() {

	var err error
	b, err := ioutil.ReadFile(CONFIG_PATH)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(b, Cfg); err != nil {
		panic(err)
	}

	if Cfg.Produce.Port == ""{
		Cfg.Produce.Port = PRODUCE_PORT
	}

}

