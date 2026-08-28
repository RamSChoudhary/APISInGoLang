package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoLang/APIs/internal/config"
)

func main() {

	fmt.Println("First apis in GoLang")

	//load config
	cfg := config.MustLoad()

	//database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to go lang apis"))
	})
	// set up server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Failed to start server")
	}

	fmt.Println("server started")
}
