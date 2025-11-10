package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/r2pq-suite/r2pq-devnet/internal/ledger"
)

type Server struct {
	state *ledger.State
	mux   *http.ServeMux
}

func NewServer(st *ledger.State) *Server {
	s := &Server{
		state: st,
		mux:   http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) Router() *http.ServeMux { return s.mux }

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.handleHealth)
	s.mux.HandleFunc("/version", s.handleVersion)
	s.mux.HandleFunc("/account/", s.handleAccount) // /account/{addr}
	s.mux.HandleFunc("/tx", s.handleTx)            // POST
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":      true,
		"ts":      time.Now().UTC().Format(time.RFC3339),
		"service": "r2pq-devnet",
	})
}

func (s *Server) handleVersion(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"version": s.state.Version()})
}

func (s *Server) handleAccount(w http.ResponseWriter, r *http.Request) {
	// path: /account/{addr}
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/account/"), "/")
	addr := strings.TrimSpace(parts[0])
	if addr == "" {
		http.Error(w, "missing address", http.StatusBadRequest)
		return
	}
	if acc, ok := s.state.GetAccount(addr); ok {
		writeJSON(w, http.StatusOK, map[string]any{
			"address": addr,
			"account": acc,
		})
		return
	}
	http.Error(w, "account not found", http.StatusNotFound)
}

func (s *Server) handleTx(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}
	var tx ledger.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "bad json: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.state.ApplyTx(tx); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Dummy hash (time-based) until real hashing is wired
	h := fmt.Sprintf("0x%X", time.Now().UnixNano())
	writeJSON(w, http.StatusOK, map[string]string{"txHash": h, "status": "accepted"})
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
