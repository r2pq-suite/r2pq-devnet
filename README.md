> **Historical 2025 R2PQ prototype lineage.** This repository preserves
> earlier exploratory work. It does not represent the architecture or
> capabilities of the current R2PQ cryptographic-assessment research system.
> Documentation below records historical implementation, limitations,
> and intentions; planned features must not be treated as delivered.
> Current public scope and evidence boundaries: [r2pq.dev](https://r2pq.dev).

## Historical prototype documentation — 2025

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
