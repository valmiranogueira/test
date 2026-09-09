FROM golang:1.26.3-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
COPY my_company/main.go ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /app .

FROM scratch

COPY --from=builder /app /app

ENTRYPOINT ["/app"]
