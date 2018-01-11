package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gotoolkit/cms/pkg/database"
	"github.com/gotoolkit/cms/pkg/handlers"
	"github.com/gotoolkit/cms/pkg/version"
)

func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	dbSource := os.Getenv("MYSQL_DATABASE")
	if port == "" {
		log.Fatal("Mysql database source is not set.")
	}

	err := database.Setup(dbSource)
	if err != nil {
		log.Fatalf("Error open mysql database connection, %s", err)
	}

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	
	if err := database.CloseDB(); err != nil {
		log.Fatal("Database Shutdown:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
