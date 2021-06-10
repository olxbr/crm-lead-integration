FROM golang:1.16.5-alpine3.13 as builder

WORKDIR /lead-integration

COPY . .

RUN go mod download && go mod verify

RUN go build -o bin/lead-integration

FROM alpine:3.13

COPY --from=builder /lead-integration/bin/lead-integration /app/lead-integration
COPY --from=builder /lead-integration/.env /app/.env

WORKDIR /app

EXPOSE 8000

ENTRYPOINT ["/app/lead-integration"]