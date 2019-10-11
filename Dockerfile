# Build binary
FROM golang:1.12
WORKDIR /go/src/github.com/haveiss/galaxy-merchant-trading
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o galaxy_trading ./cmd/galaxy_trading

FROM scratch
WORKDIR /app
# the galaxy-merchant-trading binary:
COPY --from=0 /go/src/github.com/haveiss/galaxy-merchant-trading/galaxy_trading /app/
ENTRYPOINT ["/app/galaxy_trading"]