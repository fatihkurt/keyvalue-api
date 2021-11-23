package server

import (
	"context"
	"deliveryhero/helper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupServer() {

	router := NewRouter()

	port := helper.GetEnv("PORT", "500")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Shutting down")
	os.Exit(0)
}
