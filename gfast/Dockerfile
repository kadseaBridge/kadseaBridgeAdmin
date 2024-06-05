# 使用官方 Go 语言镜像作为构建阶段的基础镜像
FROM golang:1.20 AS build

# 设置固定的项目路径
WORKDIR /src

# 将 go.mod 和 go.sum 复制到工作目录中
COPY go.mod go.sum ./

# 设置代理（如果需要的话，可以切换到 goproxy.cn）
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod download

# 将项目文件复制到工作目录中
COPY . .

# 构建应用程序
RUN go build -o /src/adminbackend ./main.go

CMD ["./adminbackend"]


