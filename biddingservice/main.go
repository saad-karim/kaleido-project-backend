package main

import (
	"fmt"
	"net/http"
	"os"

	gmux "github.com/gorilla/mux"
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/config"
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/server"
	"github.com/saad-karim/kaleido-project-backend/biddingservice/pkg/service"
)

func main() {
	db, err := service.DBOpen("postgres://ehmebenp:ZpzdWGX6WUCmT20psUP-jP4jmOAZb642@ziggy.db.elephantsql.com:5432/ehmebenp")
	if err != nil {
		fmt.Println("DB ERROR: ", err)
		os.Exit(1)
	}

	cfg := config.APIGateway()
	cfg.KaleidoAuthPassword = os.Getenv("KaleidoAuthPassword")
	cfg.KaleidoAuthUsername = os.Getenv("KaleidoAuthUsername")

	auctionLifecycleService := &service.Bidding{
		Config: cfg,
		Client: &http.Client{},
		DB:     db,
	}

	router := &gmux.Router{}
	httpServer := &http.Server{
		Addr:    "0.0.0.0:3100",
		Handler: router,
	}

	server := server.Server{
		Service:    auctionLifecycleService,
		HTTPServer: httpServer,
		Router:     router,
	}

	fmt.Println("Server listing on addr: ", httpServer.Addr)

	server.RegisterEndpoints()
	err = server.Start()
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
}
