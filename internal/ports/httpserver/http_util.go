package httpserver

import (
	"context"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/adapter/logging"
	"go-blog-ddd/internal/adapter/utils"
	"net/http"
	"strings"
	"time"
)

var logger *log.Logger

func init() {
	logger = logging.WithPrefix("http server")
	if config.Cfg.Build.Release {
		gin.SetMode(gin.ReleaseMode)
	}
}

func RunHttpServer(addr string, serverInterface ServerInterface) {

	addr = parseAddr(addr)
	//router := chi.NewRouter()

	engine := gin.New()

	setupMiddleware(engine)

	RegisterHandlers(engine, serverInterface)

	serve := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	utils.OnShutdown(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := serve.Shutdown(ctx); err != nil {
			return fmt.Errorf("Error during server shutdown: %v\n", err)
		}
		return nil
	})
	logger.Infof("running http://%s", lipgloss.NewStyle().Underline(true).Foreground(lipgloss.Color("#7CFB74")).Render(addr))
	_ = serve.ListenAndServe()
}

func parseAddr(addr string) string {
	if strings.HasPrefix(addr, ":") {
		addr = "127.0.0.1" + addr
	}
	return addr
}

func setupMiddleware(engine *gin.Engine) {
	engine.Use(
		gin.Recovery(),
	)
	if gin.Mode() == gin.DebugMode {
		engine.Use(gin.Logger())
	}
}
