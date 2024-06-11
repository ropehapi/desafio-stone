FROM golang:1.22.3 as BUILDER

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o desafio-stone cmd/server/main.go cmd/server/wire_gen.go

FROM golang:1.22.3-alpine as RUNNER

RUN apk add --no-cache bash mysql-client

WORKDIR /app

COPY --from=BUILDER /app/desafio-stone .
COPY cmd/server/.env /app/.env
COPY wait-for.sh /app/wait-for.sh

COPY --from=BUILDER /app/desafio-stone .

EXPOSE 8080

CMD ["/app/wait-for.sh", "mysql", "./desafio-stone"]