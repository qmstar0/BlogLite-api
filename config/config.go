package config

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"go-blog-ddd/internal/adapter/logging"
)

var Cfg Config

var logger = logging.WithPrefix("config")

func init() {
	logger.SetLevel(log.DebugLevel)
}

type Config struct {
	UploadFile  UploadFile  `toml:"upload_file" mapstructure:"upload_file"`
	Resource    Resource    `toml:"resource" mapstructure:"resource"`
	HttpRequest HttpRequest `toml:"http_request" mapstructure:"http_request"`
	Postgre     Postgre     `toml:"postgre" mapstructure:"postgre"`
	App         App         `toml:"app" mapstructure:"app"`
	Build       Build       `toml:"build" mapstructure:"build"`
}

type UploadFile struct {
	MaxFileSize int64 `toml:"max_file_size" mapstructure:"max_file_size"`
}
type Resource struct {
	Static struct {
		PostFromPath string `toml:"post_from_path" mapstructure:"post_from_path"`
	} `toml:"static" mapstruceure:"static"`
}

type HttpRequest struct {
	Post struct {
		DefaultLimit   int `toml:"default_limit" mapstructure:"default_limit"`
		RecentPostsNum int `toml:"recent_posts_num" mapstructure:"recent_posts_num"`
	} `toml:"post" mapstructure:"post"`
}

type Postgre struct {
	Addr     string `toml:"addr" mapstructure:"addr"`
	User     string `toml:"user" mapstructure:"user"`
	Password string `toml:"password" mapstructure:"password"`
	Database string `toml:"database" mapstructure:"database"`
}

type App struct {
	Addr string `toml:"addr" mapstructure:"addr"`
}

type Build struct {
	Release bool `toml:"release" mapstructure:"release"`
}

func Init() {
	parser := viper.NewWithOptions()
	parser.AutomaticEnv()
	parser.SetConfigName("config")
	parser.SetConfigType("toml")
	parser.AddConfigPath("config/")

	err := parser.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = parser.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
	if !Cfg.Build.Release {
		logging.Logger.SetLevel(log.DebugLevel)
		logger.Debug(parser.AllSettings())
		logger.Debug(Cfg)
	}
}
