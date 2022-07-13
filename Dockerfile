FROM golang:1.17-buster AS builder
WORKDIR /go/src/app
RUN apt-get update
COPY . .
WORKDIR /go/src/app/cmd/wallet
RUN go build -o /wallet

FROM debian:buster
RUN apt-get update
COPY ./config.json .
COPY --from=builder /wallet /wallet
CMD ./wallet