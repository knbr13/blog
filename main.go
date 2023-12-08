package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
