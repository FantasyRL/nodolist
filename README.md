# NodoList

**nodolist-demo** is a small todolist backend using golang

## deploy by docker(net=host)

(使用前请先关闭本机的mysql与redis服务)

`快速启动`
```bash
docker-compose up -d # 启动相关容器
docker build -t nodolist . # 构建镜像
docker run -d --net=host nodolist go run nodolist # 运行程序
```