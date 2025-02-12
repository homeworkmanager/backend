FROM golang:1.23.1 AS build
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o scheduleUpdater ./cmd/workers/schedule_updater

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/scheduleUpdater .

CMD ["./scheduleUpdater"]