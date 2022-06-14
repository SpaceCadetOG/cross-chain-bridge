package main

import (
	"cross-chain-bridge/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// API to interact with the functions we wrote through HTTP endpoints. To do this we are going to use gorilla/mux.

func main() {
	// connect

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	r.Handle("/api/v1/eth/{module}", handler.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))

}
