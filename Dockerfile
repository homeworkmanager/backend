FROM golang:1.23.1
ENV DEBIAN_FRONTEND=noninteractive
WORKDIR /app

ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_DB_NAME
ARG POSTGRES_HOST
ARG POSTGRES_PORT

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN chmod +x /app/entrypoint.sh

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o homeworkManager ./cmd/service





ENTRYPOINT ["./entrypoint.sh"]
CMD ["./homeworkManager"]