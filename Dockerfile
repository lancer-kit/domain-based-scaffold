# Compile stage
FROM golang:alpine AS build-env
ENV GO111MODULE=on

WORKDIR /service
ADD . .
RUN go mod download && go build -o /app .

# Final stage
FROM alpine:3.7

# Port 8080 belongs to our application
EXPOSE 8080

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache ca-certificates bash

WORKDIR /

COPY --from=build-env /app /

# Run delve
CMD ["/app", "serve"]
