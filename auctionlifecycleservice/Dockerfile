FROM golang:1.15-alpine as builder

WORKDIR /go/src
COPY . /go/src

RUN go build -o /bin/auctionlifecycleservice

FROM alpine:3.13.0
COPY --from=builder /bin/auctionlifecycleservice /usr/local/bin
EXPOSE 3100
CMD [ "/usr/local/bin/auctionlifecycleservice" ]