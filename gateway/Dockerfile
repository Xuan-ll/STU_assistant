# 使用官方的 Go 语言镜像作为基础镜像
FROM golang:alpine AS builder

# 设置工作目录为 /app
WORKDIR /app

# 将当前项目目录的所有文件拷贝到容器的 /app 目录中
COPY . .
COPY ../idl /app/idl

# 设置 Go 模块代理为 https://goproxy.cn（在中国加速模块下载），并下载项目的依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download

# 编译 Go 项目，生成可执行文件
RUN go build -o /app/bin/gateway /app/cmd/main.go

# 使用一个更小的基础镜像（Alpine）来运行应用程序
FROM alpine

# 设置工作目录为 /app
WORKDIR /app

# 从编译阶段的镜像中拷贝编译后的二进制文件到运行镜像中
COPY --from=builder /app /app
# 暴露容器的4000端口，用于外部访问
EXPOSE 4000
CMD ["bin/gateway"]