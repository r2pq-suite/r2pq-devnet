r2pq-devnet

A tiny **mock R2PQ network** for local development and testing.  
Zero external deps, safe to embed in CI. Exposes a minimal JSON/HTTP API.

## Endpoints

- `GET /health` → `{ ok, ts, service }`
- `GET /version` → `{ version }`
- `GET /account/{addr}` → `{ address, account:{ balance, nonce } }`
- `POST /tx` with JSON:
  ```json
  {
    "from": "r2pq1faucetxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "to":   "r2pq1receiverxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "amount": 100,
    "nonce": 1,
    "sig": ""
  }

→ { "txHash": "0x...", "status": "accepted" }

> NOTE: Nonce must be sender_nonce + 1. Balances are tracked in-memory only.
