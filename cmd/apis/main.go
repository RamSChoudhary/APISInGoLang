package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GoLang/APIs/internal/config"
	"github.com/GoLang/APIs/internal/http/handlers/apis"
)

func main() {

	fmt.Println("First apis in GoLang")

	//load config
	cfg := config.MustLoad()

	//database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /apis/sample", apis.New())

	// set up server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("server started")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shoutdown successfully")
}
