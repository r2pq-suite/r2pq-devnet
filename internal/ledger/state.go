package ledger

import (
	"errors"
	"sync"
)

const version = "devnet-0.1.0"

type Account struct {
	Balance uint64 `json:"balance"`
	Nonce   uint64 `json:"nonce"`
}

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
	Nonce  uint64 `json:"nonce"`
	// Placeholder for future PQ signature bytes/hex
	Sig string `json:"sig,omitempty"`
}

type State struct {
	mu       sync.RWMutex
	accounts map[string]*Account
}

func NewState(genesis map[string]Account) *State {
	accts := make(map[string]*Account, len(genesis))
	for k, v := range genesis {
		// copy to pointers
		a := v
		accts[k] = &a
	}
	return &State{accounts: accts}
}

func (s *State) Version() string { return version }

func (s *State) GetAccount(addr string) (Account, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.accounts[addr]
	if !ok {
		return Account{}, false
	}
	return *a, true
}

var (
	ErrUnknownSender   = errors.New("unknown sender")
	ErrInsufficientBal = errors.New("insufficient balance")
	ErrBadNonce        = errors.New("bad nonce")
)

// ApplyTx moves balance and bumps the sender nonce.
// Super-simple checks; no fees, no signatures yet.
func (s *State) ApplyTx(tx Transaction) error {
	if tx.From == "" || tx.To == "" {
		return errors.New("missing from/to")
	}
	if tx.Amount == 0 {
		return errors.New("amount must be > 0")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	from, ok := s.accounts[tx.From]
	if !ok {
		return ErrUnknownSender
	}
	to, ok := s.accounts[tx.To]
	if !ok {
		// auto-create receiver
		to = &Account{}
		s.accounts[tx.To] = to
	}

	// very simple nonce rule
	expected := from.Nonce + 1
	if tx.Nonce != expected {
		return ErrBadNonce
	}

	if from.Balance < tx.Amount {
		return ErrInsufficientBal
	}

	from.Balance -= tx.Amount
	from.Nonce++
	to.Balance += tx.Amount
	return nil
}
