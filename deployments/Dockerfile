# base stage
FROM golang:alpine AS base

WORKDIR /app

COPY go.mod . 
COPY go.sum .

# build stage
FROM golang:alpine AS builder

WORKDIR /app

COPY --from=base /app ./
RUN go mod download
COPY . .

RUN go build -o amartha-loan-system main.go

# run stage
FROM alpine:latest AS runner

WORKDIR /app

RUN apk add --no-cache tzdata
COPY --from=builder /app/amartha-loan-system ./

CMD ["./amartha-loan-system"]