package config

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Application struct {
	ReadTimeout   int    `yaml:"read_timeout"`
	WriterTimeout int    `yaml:"writer_timeout"`
	Host          string `yaml:"host"`
	Port          int64  `yaml:"port"`
	Name          string `yaml:"name"`
	JwtSecret     string `yaml:"jwt_secret"`
	Mode          string `yaml:"Mode"`
	DemoMsg       string `yaml:"DemoMsg"`
	EnableDP      bool   `yaml:"enable_dp"`
}
type Database struct {
	Driver          string `yaml:"driver"`
	Source          string `yaml:"source"`
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time"`
	ConnMaxLifeTime int    `yaml:"conn_max_life_time"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
}
type Jwt struct {
	Secret  string `yaml:"secret"`
	Timeout int64  `yaml:"timeout"`
}

type Logger struct {
	Type      string `yaml:"type"`
	Path      string `yaml:"path"`
	Level     string `yaml:"level"`
	Stdout    string `yaml:"stdout"`
	EnabledDB bool   `yaml:"enabled_db"`
	Cap       uint   `yaml:"cap"`
}
type Config struct {
	Application Application `yaml:"application"`
	Database    Database    `yaml:"database"`
	Jwt         Jwt         `yaml:"jwt"`
	Logger      Logger      `yaml:"logger"`
}

var Cfg *Config

func Init() (err error) {
	// 1. 读取配置文件内容并解析为 Config 对象
	Cfg, err = loadConfig("config/settings.yml")
	if err != nil {
		log.Fatal("loadConfig() failed")
		return err
	}
	return nil
}

func loadConfig(filename string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("ioutil.ReadFile() failed ")
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
