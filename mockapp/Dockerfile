FROM golang:1.22-alpine AS builder

WORKDIR /app 

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.* ./
RUN go mod download && go mod verify

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/benchmarkapp ./cmd/benchmarkapp.go

FROM alpine:3.14.10

EXPOSE 8080

COPY --from=builder /app/bin/benchmarkapp .

# Include dumb-init to avoid application starting at PID 1
# See https://github.com/Yelp/dumb-init
RUN apk add dumb-init
ENTRYPOINT [ "/usr/bin/dumb-init", "--"]

# Run binary
CMD [ "/benchmarkapp" ]
