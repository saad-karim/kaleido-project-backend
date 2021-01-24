package main

import (
	"fmt"
	"net/http"
	"os"

	gmux "github.com/gorilla/mux"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/config"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/server"
	"github.com/saad-karim/kaleido-project-backend/auctionlifecycleservice/pkg/service"
)

func main() {
	db, err := service.DBOpen("postgres://ehmebenp:ZpzdWGX6WUCmT20psUP-jP4jmOAZb642@ziggy.db.elephantsql.com:5432/ehmebenp")
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	auctionLifecycleService := &service.AuctionLifecycle{
		Config: config.APIGateway(),
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
