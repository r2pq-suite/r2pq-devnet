package main

import (
	"log"
	"net/http"

	"github.com/r2pq-suite/r2pq-devnet/internal/ledger"
	"github.com/r2pq-suite/r2pq-devnet/internal/mock"
	"github.com/r2pq-suite/r2pq-devnet/internal/rpc"
)

func main() {
	// Initialize state with a tiny genesis
	st := ledger.NewState(mock.Genesis())

	srv := rpc.NewServer(st)

	addr := ":7878"
	log.Printf("r2pq-devnet listening on %s", addr)
	if err := http.ListenAndServe(addr, srv.Router()); err != nil {
		log.Fatal(err)
	}
}
