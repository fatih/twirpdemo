package main

import (
	"log"
	"net/http"

	"github.com/fatih/twirpdemo/internal/calculator"
	"github.com/fatih/twirpdemo/internal/rpc"
)

func main() {
	// create the business logic satisfying the generated Twirp Calculator
	// interface.
	svc := calculator.NewCalculatorService()

	// server satisfies the http.Handler interface.
	server := rpc.NewCalculatorServer(svc)

	port := "8080"
	log.Println("server started at port", port)
	http.ListenAndServe(":"+port, server)
}
