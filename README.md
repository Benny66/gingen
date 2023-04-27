# gingen

初始化项目并编译
```
go mod tidy
go build
```
执行并在当前目录新增ginServer框架内容
```
go run main.go --mod testServer
cd testServer
go mod tidy
//配置.env中 端口、服务地址、数据库等信息
go run main.go 
```

执行命令即可在当前目录生成[ginServer](https://github.com/Benny66/ginServer.git)框架内容，testServer为项目


```
mkdir testServer
cd testServer
./gingen init --mod testServer
```