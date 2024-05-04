package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	UploadFile UploadFile `toml:"upload_file" mapstructure:"upload_file"`
	Resource   Resource   `toml:"resource" mapstructure:"resource"`
	Request    Request    `toml:"request" mapstructure:"request"`
	Postgre    Postgre    `toml:"postgre" mapstructure:"postgre"`
}

type UploadFile struct {
	MaxFileSize int64 `toml:"max_file_size" mapstructure:"max_file_size"`
}
type Resource struct {
	Static struct {
		PostFromPath string `toml:"post_from_path" mapstructure:"post_from_path"`
	} `toml:"static" mapstruceure:"static"`
}

type Request struct {
	Post struct {
		DefaultLimit int `toml:"default_limit" mapstructure:"default_limit"`
	} `toml:"post" mapstructure:"post"`
}

type Postgre struct {
	Addr     string `toml:"addr" mapstructure:"addr"`
	User     string `toml:"user" mapstructure:"user"`
	Password string `toml:"password" mapstructure:"password"`
	Database string `toml:"database" mapstructure:"database"`
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}

	fmt.Println(viper.AllSettings())
	fmt.Println(Conf)
}
