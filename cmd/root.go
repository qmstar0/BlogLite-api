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
	"os"
)

var (
	configPath string
	logger     *log.Logger
)

const defaultConfigPath = "config.toml"

var rootCmd = &cobra.Command{
	Use:   "domain-api",
	Short: "Start the server",
	Long:  `Start the server on the specified flagport.`,
	//运行前hooks
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//初始化配置
		config.Init(configPath, cmd.Flags())
		//初始化数据库
		dbCloseFn := postgresql.Init(config.Cfg.DatabaseDNS)
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
		logger = logging.Logger().WithPrefix("server")
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
	rootCmd.Flags().StringVar(&configPath, "config", defaultConfigPath, "Set configuration file path")
	rootCmd.Flags().IntP("port", "p", 5000, "Set the admintoken running flagport")
	rootCmd.Flags().Bool("release", false, "Set mode")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err)
		shutdown.Exit(1)
	}
	shutdown.WaitCtrlC()
}
