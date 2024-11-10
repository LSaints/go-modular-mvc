FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
COPY .env ./

RUN go mod download

COPY . ./

RUN chmod 644 .env && \
    CGO_ENABLED=0 GOOS=linux go build -o /build/main ./cmd/app/main.go

FROM debian:bullseye-slim AS release-stage

WORKDIR /go-app

COPY --from=build-stage /app/.env ./.env
COPY --from=build-stage /app/internal ./internal/
COPY --from=build-stage /build/main ./main

RUN chmod 644 .env

EXPOSE 8080

ENTRYPOINT ["/go-app/main"]
