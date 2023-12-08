package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		log.Panicln("caught signal", map[string]string{
			"signal": s.String(),
		})
		os.Exit(0)
	}()

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	err := srv.Shutdown(context.Background())
	if err != nil {
		log.Fatal("error shutting down: ", err)
	}
	time.Sleep(time.Second * 2)
}
