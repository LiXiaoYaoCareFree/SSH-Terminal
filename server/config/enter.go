package config

import "fmt"

type Config struct {
	System System `yaml:"system"`
	Ssh    Ssh    `yaml:"ssh"`
}

type System struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Ssh struct {
	DestIP   string `yaml:"destIP"`   // 目标ip
	DestPort int    `yaml:"destPort"` // 目标端口
	User     string `yaml:"user"`     // 用户名
	Pwd      string `yaml:"pwd"`      // 密码
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
