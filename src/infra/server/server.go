package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hpaes/go-api-final-project/src/infra/router"
)

type (
	Port int64

	ServerConfig interface {
		Listen(ctx context.Context, wg *sync.WaitGroup)
	}

	webServer struct {
		r       router.GinRouter
		port    int64
		ctx     context.Context
		timeout time.Duration
	}
)

func NewWebServer(r router.GinRouter, port int64, timeout time.Duration) *webServer {
	return &webServer{
		r:       r,
		port:    port,
		timeout: timeout,
	}
}

func (ws *webServer) Listen(ctx context.Context, wg *sync.WaitGroup) {
	gin.SetMode(gin.ReleaseMode)

	ws.r.SetAppHandlers()

	ctx, cancel := context.WithCancel(ctx)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", ws.port),
		Handler:      ws.r.GetRouter(),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-stop
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error: %v\n", err)
		}
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutDownCtx, shutDownCancel := context.WithTimeout(context.Background(), ws.timeout)
		defer shutDownCancel()

		if err := server.Shutdown(shutDownCtx); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Server successfully stopped")
		}
	}()
}
