FROM golang:latest AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -ldflags "-s -w" -o backend .

FROM scratch as runner

COPY --from=builder /build/backend /
COPY --from=builder /build/config/settings.yml /config/
COPY --from=builder /build/config/db.sql /config/

EXPOSE 8080

ENTRYPOINT ["/backend"]
