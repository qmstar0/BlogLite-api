package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	Port    int  `toml:"port" mapstructure:"port"`
	Release bool `toml:"release" mapstructure:"release"`

	UploadFile  UploadFile  `toml:"upload_file" mapstructure:"upload_file"`
	HttpRequest HttpRequest `toml:"http_request" mapstructure:"http_request"`
	Postgre     Postgre     `toml:"postgre" mapstructure:"postgre"`
	JWTAuth     JWTAuth     `toml:"jwt_auth" mapstructure:"jwt_auth"`
}

type UploadFile struct {
	MaxFileSize int64 `toml:"max_file_size" mapstructure:"max_file_size"`
}

type HttpRequest struct {
	Post struct {
		DefaultLimit   int `toml:"default_limit" mapstructure:"default_limit"`
		RecentPostsNum int `toml:"recent_posts_num" mapstructure:"recent_posts_num"`
	} `toml:"post" mapstructure:"post"`
}

type JWTAuth struct {
	AuthKey       string   `toml:"authkey" mapstructure:"authkey"`
	Subject       string   `toml:"subject" mapstructure:"subject"`
	Issuer        string   `toml:"issuer" mapstructure:"issuer"`
	Audience      []string `toml:"audience" mapstructure:"audience"`
	AuthTokenLife int      `toml:"authtoken_left" mapstructure:"authtoken_left"`
}

type Postgre struct {
	Addr     string `toml:"addr" mapstructure:"addr"`
	User     string `toml:"user" mapstructure:"user"`
	Password string `toml:"password" mapstructure:"password"`
	Database string `toml:"database" mapstructure:"database"`
}

func Init(configFile string, flags *pflag.FlagSet) {
	if flags != nil {
		_ = viper.BindPFlags(flags)
	}
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
}

func AllSettings() map[string]any {
	return viper.AllSettings()
}
