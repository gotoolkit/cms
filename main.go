package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gotoolkit/cms/pkg/database"
	"github.com/gotoolkit/cms/pkg/handlers"
	"github.com/gotoolkit/cms/pkg/teleport"
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
	if dbSource == "" {
		log.Fatal("Mysql database source is not set.")
	}

	err := database.Setup(dbSource)
	if err != nil {
		log.Fatalf("Error open mysql database connection, %s", err)
	}

	teleSource := os.Getenv("TELEGRAM_HORN_URL")
	if teleSource == "" {
		log.Fatal("Telegram notification source is not set.")
	}

	teleport.Setup(teleSource)

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
	database.CloseDB()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
