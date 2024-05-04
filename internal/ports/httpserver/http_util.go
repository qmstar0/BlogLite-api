package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog-ddd/internal/adapter/utils/shutdown"
	"net/http"
	"strings"
	"time"
)

func RunHttpServer(addr string, serverInterface ServerInterface) {
	addr = parseAddr(addr)
	//router := chi.NewRouter()

	engine := gin.New()

	setupMiddleware(engine)

	RegisterHandlers(engine, serverInterface)

	printStartInfo(addr, engine)

	serve := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	shutdown.OnShutdown(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := serve.Shutdown(ctx); err != nil {
			return fmt.Errorf("Error during server shutdown: %v\n", err)
		}
		return nil
	})

	_ = serve.ListenAndServe()
}

func parseAddr(addr string) string {
	if strings.HasPrefix(addr, ":") {
		addr = "127.0.0.1" + addr
	}
	return addr
}

func printStartInfo(addr string, router *gin.Engine) {
	//_ = chi.Walk(router, func(method string, route string, decorator http.Handler, _ ...func(http.Handler) http.Handler) error {
	//	fmt.Printf("\033[32m%-10s\033[m - \033[1;4m%-10s\033[m\n", fmt.Sprintf("[%s]", method), route)
	//	return nil
	//})

	for _, info := range router.Routes() {
		fmt.Printf("\033[32m%-10s\033[m - \033[1;4m%-10s\033[m\n", fmt.Sprintf("[%s]", info.Method), info.Path)
	}

	fmt.Printf("\n\033[1mHttpserver Starts Running: \033[m\033[1;4;32m%s\033[m\n", addr)
}

func setupMiddleware(engine *gin.Engine) {
	engine.Use(
		gin.Recovery(),
		//gin.Logger(),
	)
}
