# 该镜像用于编译web程序
FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY="https://goproxy.cn,direct" \
    GO111MODULE=on

WORKDIR /build
COPY ["main.go","go.mod","go.sum","simkai.ttf","./"]
COPY ["/conf","./"]
COPY . .

RUN go mod download 
RUN go build -ldflags="-s -w" -o /app/route_planner


#该镜像用于运行web程序
FROM alpine

MAINTAINER np<20002515@mail.ecust.edu.cn>

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/route_planner /app/
COPY --from=builder /build/conf/application.yaml /app/conf/

EXPOSE 8080

CMD ["./route_planner"]
