## BUILDER

FROM golang:1.15-alpine AS builder

WORKDIR /app
ADD . /app

RUN go build -o server

## RUNNER

FROM alpine:3.12

WORKDIR /app
COPY --from=builder /app/server .
ADD config.toml .

CMD [ "/app/server" ]
