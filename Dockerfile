FROM golang:1.21 AS builder

MAINTAINER CodeLine "1145938682@qq.com"
RUN mkdir -p /home/work/app/go-cloud-native
WORKDIR /home/work/app/go-cloud-native
COPY . .
RUN echo go-cloud-native
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o build/package/main ./cmd/server/main.go


FROM centos:7.3
MAINTAINER go-cloud-native "1145938682@qq.com"
WORKDIR /home/work/app/go-cloud-native
COPY --from=builder /home/work/app/go-cloud-native/build/package .

USER root

CMD ["sh", "-c", "/home/work/app/go-cloud-native/main", "-c","/home/work/app/go-cloud-native/configs/local.toml"]




