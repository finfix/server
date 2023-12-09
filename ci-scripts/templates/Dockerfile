# syntax=docker/dockerfile:1
FROM golang:1.20 AS builder
RUN apt install ca-certificates
ADD . /go
WORKDIR /go/@{project_name}
RUN go generate ./...
RUN go build app/cmd/main.go

FROM ubuntu:latest
WORKDIR /@{project_name}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/@{project_name}/main .
EXPOSE @{service_port}
USER root:root
CMD ["./main"]
