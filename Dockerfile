FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY my_company/go.mod ./
COPY my_company/go.sum ./
COPY my_company/main.go ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /app .

FROM scratch

COPY --from=builder /app /app

ENTRYPOINT ["/app"]
