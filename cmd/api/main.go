package main

import (
	"CatalogItems/internal/products/app"
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	var err error

	a, err := app.NewApp()

	if err != nil {
		log.Fatal(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Server started on :8080")
		if err := a.Run(); err != nil {
			log.Println("server stopped:", err)
		}
	}()
	<-stop
	log.Println("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.Shutdown(ctx); err != nil {
		log.Println("server shutdown:", err)
	}

	log.Println("Server stopped gracefully")
}
