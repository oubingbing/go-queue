## 部署流程

### 一、安装golang环境，需要1.13以上的版本
	wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
	tar -C /usr/local/ -xvf go1.14.linux-amd64.tar.gz
	vim /etc/profile
	export GOROOT=/usr/local/go
	export PATH=$PATH:/usr/local/go/bin
	export GOPATH=/data/golang
	
	go version

### 二、拉取代码
发布地址： (git pull)
发布分支： master
发布备注：

### 三、设置mod环境变量
	vim /etc/profile
	GOPROXY=https://goproxy.io
	GO111MODULE=on

### 四、打包
在项目目录下执行一下命令

	set GOOS=linux
	set GOARCH=amd64
	go build

### 五、启动服务
先给生成的二级制文件scoket读写的权限
然后执行以下命令启动服务

`nohup ./project &`
	
完成部署

### 配置参数

```
[redis]
HOST = localhost:6379
PASSWORD =
DB = 0

# Python打包工具回调写入的队列名称
REDIS_DB_CALLBACK_KEY = package_finish_list

[mysql]
DB_DRIVER=
DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=

[socket]
SOCKET_URL=

[log_path]
LOG_PATH=
```
