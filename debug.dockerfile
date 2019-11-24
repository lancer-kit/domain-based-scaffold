# Compile stage
FROM golang:alpine AS build-env
ENV CGO_ENABLED 0
ENV GO111MODULE=on

WORKDIR /service

ADD . .
RUN go mod download
# The -gcflags "all=-N -l" flag helps us get a better debug experience
RUN go build -gcflags "all=-N -l" -o -o /app .

# Compile Delve
RUN apk add --no-cache git
RUN go get github.com/derekparker/delve/cmd/dlv

# Final stage
FROM alpine:3.7

# Port 8080 belongs to our application, 40000 belongs to Delve
EXPOSE 8080 40000

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat ca-certificates bash

WORKDIR /

COPY --from=build-env /app /
COPY --from=build-env /go/bin/dlv /

# Run delve
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/app"]
