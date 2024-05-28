package httpserver

import (
	"context"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/nightsky-api/config"
	"github.com/qmstar0/nightsky-api/internal/apps/service"
	"github.com/qmstar0/nightsky-api/pkg/logging"
	"net/http"
	"time"
)

const DefaultHttpServerPort = 5000

type HttpServeLauncher struct {
	server *http.Server
	logger *log.Logger
}

func NewHttpServeLauncher(serverInterface ServerInterface, service *service.AdminAuthenticationService) *HttpServeLauncher {
	if config.Cfg.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	setupMiddleware(engine)
	RegisterHandlersWithOptions(engine, serverInterface, GinServerOptions{
		BaseURL:      "/api",
		Middlewares:  []MiddlewareFunc{AuthMiddleware(service)},
		ErrorHandler: nil,
	})

	serve := &http.Server{Handler: engine}
	return &HttpServeLauncher{
		server: serve,
		logger: logging.WithPrefix("http server"),
	}
}

func (h HttpServeLauncher) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.server.Shutdown(ctx); err != nil {
		h.logger.Error("Error during server shutdown: %v\n", err)
	}
}

func (h HttpServeLauncher) Run(port int) {
	if port <= 0 {
		port = DefaultHttpServerPort
	}
	h.server.Addr = h.spliceAddr(port)
	h.logger.Printf("running http://%s", lipgloss.NewStyle().Underline(true).Foreground(lipgloss.Color("#7CFB74")).Render(h.server.Addr))
	_ = h.server.ListenAndServe()
}

func (h HttpServeLauncher) spliceAddr(port int) string {
	host := "127.0.0.1"
	if config.Cfg.Release {
		host = "0.0.0.0"
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func setupMiddleware(engine *gin.Engine) {
	engine.Use(
		gin.Recovery(),
	)
	if gin.Mode() == gin.DebugMode {
		engine.Use(gin.Logger())
	}
}
