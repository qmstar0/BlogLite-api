package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/qmstar0/domain/config"
	"github.com/qmstar0/domain/internal/apps"
	"github.com/qmstar0/domain/internal/apps/service"
	"github.com/qmstar0/domain/pkg/logging"
	"github.com/qmstar0/domain/pkg/postgresql"
	"github.com/qmstar0/domain/ports/httpserver"
	"github.com/qmstar0/shutdown"
	"github.com/spf13/cobra"
)

var (
	configPath string
	logger     *log.Logger
)

const defaultConfigPath = "config.toml"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server on the specified flagport.`,
	//运行前hooks
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init(configPath, cmd.Flags())
		dbCloseFn := postgresql.Init(
			config.Cfg.Postgre.Addr,
			config.Cfg.Postgre.User,
			config.Cfg.Postgre.Password,
			config.Cfg.Postgre.Database,
		)
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
		httpServer := httpserver.NewHttpServer(apps.NewDomainApp(), service)
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
