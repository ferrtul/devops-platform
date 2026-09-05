FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o /app/server .

    
FROM alpine:3.22

RUN addgroup -S appgroup && \
    adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/server .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./server"]