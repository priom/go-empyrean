# Build Geth in a stock Go builder container
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-ethereum
RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/
COPY --from=builder /go-ethereum/config.toml /usr/local/bin/
COPY --from=builder /go-ethereum/initShyftGeth_docker.sh /usr/local/bin/
COPY --from=builder /go-ethereum/resetShyftGeth_docker.sh /usr/local/bin/
COPY --from=builder /go-ethereum/startShyftGeth_docker.sh /usr/local/bin/
COPY --from=builder /go-ethereum/unlockPasswords.txt /usr/local/bin/
COPY --from=builder /go-ethereum/ShyftNetwork.json /usr/local/bin/
COPY --from=builder /go-ethereum/init_all.sh /usr/local/bin/

COPY --from=builder /go-ethereum/UTC--2018-03-18T22-35-03.111589431Z--43ec6d0942f7faef069f7f63d0384a27f529b062 /
COPY --from=builder /go-ethereum/UTC--2018-03-18T22-36-53.786508893Z--9e602164c5826ebb5a6b68e4afd9cd466043dc4a /
COPY --from=builder /go-ethereum/UTC--2018-03-18T22-37-04.212101982Z--5bd738164c61fb50eb12e227846cbaef2de965aa /
COPY --from=builder /go-ethereum/UTC--2018-03-18T22-37-12.082606217Z--c04ee4131895f1d0c294d508af65d94060aa42bb /
COPY --from=builder /go-ethereum/UTC--2018-03-18T22-37-19.650676366Z--07d899c4ac0c1725c35c5f816e60273b33a964f7 /

EXPOSE 8545 8546 30303 30303/udp 30304/udp
ENTRYPOINT ["/usr/local/bin/init_all.sh"]
