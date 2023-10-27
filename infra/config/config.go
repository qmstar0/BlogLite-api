package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	err  error
	Conf = new(Config)
)

type Config struct {
	Service  *Service  `yaml:"Service"`
	Database *Database `yaml:"Database"`
	Logger   *logger   `yaml:"Logger"`
	Mail     *Mail     `yaml:"Mail"`
	Smtp     *Smtp     `yaml:"Smtp"`
	Redis    *Redis    `yaml:"Redis"`

	Jwt     *JWT     `yaml:"Jwt"`
	System  *System  `yaml:"System"`
	Article *Article `yaml:"Article"`
	User    *User    `yaml:"User"`
	Event   *Event   `yaml:"Event"`
}

func init() {
	s := os.Getenv("BLOG_CONFIG_PATH")
	if s == "" {
		// panic("BLOG_CONFIG_PATH 没有配置")
		s = "system/config.yml"
	}
	InitDefaultConfig()
	viper.SetConfigFile(s)
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	//fmt.Printf(">>%v\n", *Conf.Jwt)
	//fmt.Printf(">>%v\n", *Conf.User)
	//fmt.Printf(">>%v\n", *Conf.Article)
	//fmt.Printf(">>%v\n", *Conf.System)
}
