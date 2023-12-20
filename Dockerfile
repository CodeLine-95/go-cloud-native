FROM golang:1.21-alpine AS builder

MAINTAINER CodeLine "1145938682@qq.com"
RUN mkdir -p /home/work/app/go-cloud-native
WORKDIR /home/work/app/go-cloud-native
COPY . .
RUN echo go-cloud-native \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download \
    && go build -o build/package/main ./cmd/server/main.go


# 使用最小镜像来运行编译后的 go 项目
FROM alpine as runner
MAINTAINER go-cloud-native "1145938682@qq.com"
WORKDIR /home/work/app/go-cloud-native
COPY --from=builder /home/work/app/go-cloud-native/build/package ./build/package
COPY --from=builder /home/work/app/go-cloud-native/configs ./configs

EXPOSE 8000

CMD ["/home/work/app/go-cloud-native/build/package/main", "-c", "/home/work/app/go-cloud-native/configs/local.toml"]




