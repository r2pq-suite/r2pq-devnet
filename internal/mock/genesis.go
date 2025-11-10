package mock

import "github.com/r2pq-suite/r2pq-devnet/internal/ledger"

func Genesis() map[string]ledger.Account {
	return map[string]ledger.Account{
		// Faucet account with initial supply
		"r2pq1faucetxxxxxxxxxxxxxxxxxxxxxxxxxxxx": {Balance: 1_000_000_000, Nonce: 0},
		// Empty example recipient
		"r2pq1receiverxxxxxxxxxxxxxxxxxxxxxxxxxx": {Balance: 0, Nonce: 0},
	}
}
