package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/apps"
	"go-blog-ddd/internal/apps/service"
	"go-blog-ddd/pkg/logging"
	"go-blog-ddd/pkg/postgresql"
	"go-blog-ddd/pkg/shutdown"
	"go-blog-ddd/ports/httpserver"
)

var (
	configPath string
	logger     *log.Logger
)

const defaultConfigPath = "config/config.toml"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server on the specified flagport.`,
	//运行前hooks
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init(configPath, cmd.Flags())
		dbCloseFn := postgresql.Init()
		shutdown.RegisterTasks(func() {
			err := dbCloseFn()
			if err != nil {
				fmt.Println(err)
			}
		})
		//初始化日志
		if config.Cfg.Release {
			logging.Init(log.WarnLevel)
		} else {
			logging.Init(log.DebugLevel)
		}

		logger = logging.Logger().WithPrefix("config")
	},
	//main run
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(config.Cfg)
		logger.Debug(config.AllSettings())

		service := service.NewAdminAuthenticationService()
		httpServer := httpserver.NewHttpServer(apps.NewDomainControl(), service)
		launcher := httpserver.NewHttpServeLauncher(httpServer, service)
		shutdown.RegisterTasks(launcher.Close)

		launcher.Run(config.Cfg.Port)
	},
}

func init() {
	serveCmd.Flags().StringVar(&configPath, "config", defaultConfigPath, "Set configuration file path")
	serveCmd.Flags().IntP("port", "p", 0, "Set the admintoken running flagport")
	serveCmd.Flags().Bool("release", false, "Set mode")
	rootCmd.AddCommand(serveCmd)
}
