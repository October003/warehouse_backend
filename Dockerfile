FROM golang:alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o warehouse .

#FROM scratch
#
#COPY --from=builder /build/warehouse /
#COPY --from=builder /build/config/settings.yml /
WORKDIR /dist

RUN cp /build/warehouse .

EXPOSE 8888

ENTRYPOINT ["/dist/warehouse"]





