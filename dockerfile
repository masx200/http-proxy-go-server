FROM docker.cnb.cool/masx200/docker_mirror/golang:1.24.4-alpine-linux-amd64 AS build


# Set the working directory
WORKDIR /build
run sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
# Install git
RUN --mount=type=cache,target=/var/cache/apk \
    apk add git ca-certificates

env GO111MODULE=on
env  GOPROXY=https://goproxy.cn

copy . .

run go mod tidy

# Build the server
# go build automatically download required module dependencies to /go/pkg/mod
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 go build -v -ldflags="-s -w "  \
    -o /bin/main ./main.go



from docker.cnb.cool/masx200/docker_mirror/http-proxy-go-server:2025-09-03-16-34-08


workdir /app

cmd     ["/app/main"]

COPY --from=build /bin/main .


run chmod +x /app/main


run sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories


RUN --mount=type=cache,target=/var/cache/apk \
    apk add git ca-certificates curl