FROM golang:1.22.3 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

#  -ldflags "-s -w" remove debug information
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o ./bin/school-be ./cmd/school-be-server/main.go

# move to separate image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/school-be /app/bin/school-be
COPY .env /app/
