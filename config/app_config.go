package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"time"
)

type AppConfig struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Sms      Sms      `yaml:"sms"`
}

type Database struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"database_name"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

type Server struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	ModeRun  string `yaml:"mode_run"`  // prod or debug
	LevelLog int    `yaml:"level_log"` // level log from 0 -> 6
}

type Sms struct {
	AccountSid    string `yaml:"account_sid"`
	AuthToken     string `yaml:"auth_token"`
	MyPhoneNumber string `yaml:"my_phone_number"`
}

func NewAppConfig(configPath string) (*AppConfig, error) {
	var appCfg AppConfig
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &appCfg)
	if err != nil {
		log.Fatalf("Cannot read file config.yaml - %s", err)
	}

	return &appCfg, nil
}

func checkConnectionIsAvailable(host, port string) error {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", host+port, timeout)
	if err != nil {
		return err
	}
	return nil
}
