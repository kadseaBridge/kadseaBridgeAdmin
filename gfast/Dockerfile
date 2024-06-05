# 使用官方 Go 语言镜像作为构建阶段的基础镜像
FROM golang:1.20 AS build

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到工作目录中
COPY go.mod go.sum ./

# 设置代理（使用 goproxy.cn）
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod download

# 将项目文件复制到工作目录中
COPY . .

# 构建应用程序
RUN go build -o myapp .

# 使用一个轻量级的基础镜像来运行应用程序
FROM alpine:latest

# 安装 ca-certificates 包，以确保 https 请求可以正常工作
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件到工作目录
COPY --from=build /app/myapp .

# 暴露应用程序运行的端口
EXPOSE 8200

# 启动应用程序
CMD ["./myapp"]
