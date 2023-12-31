ARG  BUILDER_IMAGE=golang:buster
ARG  DISTROLESS_IMAGE=gcr.io/distroless/static

############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder
EXPOSE 8080
RUN update-ca-certificates
WORKDIR /go/bin
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s' -a -o /go/bin/run cmd/main.go

ENTRYPOINT ["/go/bin/run"]