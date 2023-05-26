# gingen
命令
```
Usage:
  gingen [command]

Available Commands:
  api         create a api module
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Init a ginServer module
  model       create a model/dao module

Flags:
  -h, --help               help for gingen
      --log-level string   log level (default "info")

Use "gingen [command] --help" for more information about a command.

```

### 初始化项目并编译

```
go mod tidy
go build
//可将执行文件gingen放置go bin目录中，全局执行
cp ./gingen /usr/local/go/bin/gingen
```

### 执行并在当前目录新增ginServer框架内容
```
go run main.go --mod testServer (或者gingen --mod testServer)
cd testServer
go mod tidy
//配置.env中 端口、服务地址、数据库等信息
source .env && go run main.go 
```

执行命令即可在当前目录生成[ginServer](https://github.com/Benny66/ginServer.git)框架内容，testServer为项目

### .env配置在根目录，例如
```
export MODE=debug
export DB_HOST=127.0.0.1
export DB_PORT=33060
export DB_USERNAME=root
export DB_PASSWORD=root123456
export DB_NAME=card-task

```

### 新增api接口
```
//会同时新增api模块、service模块、schema模块
gingen api --name user
gingen api --name user --func getUserList
```

