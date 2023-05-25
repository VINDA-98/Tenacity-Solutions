
English | [简体中文](./README.md)


## 1. Getting started

```
- node version > v8.6.0
- golang version >= v1.14
- IDE recommendation: Goland
- initialization project: different versions of the database are not initialized. See synonyms at initialization https://www.gin-vue-admin.com/docs/first
- Replace the Qiniuyun public key, private key, warehouse name and default url address in the project to avoid data confusion in the test file.
```

### 1.1 server project

use `Goland` And other editing tools，open server catalogue，You can't open it. `gin-vue-admin` root directory

```bash
# clone the project
git clone https://github.com/flipped-aurora/gin-vue-admin.git

# open server catalogue
cd server

# use go mod And install the go dependency package
go generate

# Compile 
go build -o server main.go (windows the compile command is go build -o server.exe main.go )

# Run binary
./server (windows The run command is server.exe)
```

### 1.2 web project

```bash
# enter the project directory
cd web

# install dependency
npm install

# develop
npm run serve
```

### 1.2 Server

```bash
# using go.mod

# install go modules
go list (go mod tidy)

# build the server
go build
```

### 1.3 API docs auto-generation using swagger

#### 1.3.1 install swagger 

##### (1) Using VPN or outside mainland China
````
go get -u github.com/swaggo/swag/cmd/swag
````

##### (2) In mainland China

In mainland China, access to go.org/x is prohibited，we recommend [goproxy.io](https://goproxy.io/zh/) or [goproxy.cn](https://goproxy.cn)

````bash
# If you are using a version of Go 1.13 - 1.15 Need to set up manually GO111MODULE=on, The opening mode is as follows, If your Go version is 1.16 ~ Latest edition You can ignore the following step one
# Step one、Enable Go Modules Function
go env -w GO111MODULE=on 
# Step two、Configuration GOPROXY Environment variable
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# If you dislike trouble,You can use the go generate Automatically execute code before compilation, But this can't be used command line terminal of `Goland` or `Vscode` 
cd server
go generate -run "go env -w .*?"

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
````

#### 1.3.2 API docs generation

````
cd server
swag init
````