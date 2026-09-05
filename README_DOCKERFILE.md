                 BUILD STAGE
        ┌─────────────────────────┐
        │ golang:1.25-alpine      │
        │                         │
        │ /src                    │
        │ ├── go.mod              │
        │ ├── go.sum              │
        │ └── main.go             │
        │          │              │
        │          ▼              │
        │      go build           │
        │          │              │
        │          ▼              │
        │     /app/server         │
        └────────────┬────────────┘
                     │
                     │ COPY --from=builder
                     ▼
                 RUNTIME STAGE
        ┌─────────────────────────┐
        │ alpine:3.22             │
        │                         │
        │ /app                    │
        │ └── server              │
        │                         │
        │ USER appuser            │
        │ PORT 8080               │
        └────────────┬────────────┘
                     │
                     ▼
              Go application
                   :8080
                     │
                     ▼
                PostgreSQL