FROM golang:alpine AS builder
LABEL stage=gobuilder

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

ENV LANG C.UTF-8

RUN apk update --no-cache && apk add --no-cache tzdata

RUN #apt-get install -y mysql-server


WORKDIR /build

COPY . .

EXPOSE 6666

RUN go build -ldflags="-s -w" -o /app/main

FROM alpine


EXPOSE 3306

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

ENV TZ Asia/Shanghai

ENV LANG C.UTF-8

WORKDIR /app

#COPY --from=builder /build/. /app/.
COPY  . /app/.
COPY --from=builder /app/main /app/main

EXPOSE 6666

CMD ["./main"]
