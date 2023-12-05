# syntax=docker/dockerfile:1
FROM golang:1.20 AS builder

RUN apt install ca-certificates

WORKDIR /@{project_name}

COPY go.mod go.sum ./
COPY app ./app

RUN go mod download

RUN go generate ./...
RUN go build -o project app/cmd/main.go

FROM ubuntu:latest

WORKDIR /@{project_name}

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /@{project_name}/project .

EXPOSE @{service_port}

USER root:root

CMD ["./project"]
