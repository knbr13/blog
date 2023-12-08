package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
	})
	log.Fatal(srv.ListenAndServe())
}
