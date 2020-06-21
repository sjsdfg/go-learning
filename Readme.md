## Go-Learning

go 语言学习尝试

### 基础概念

1. GOROOT：Go 语言安装根目录的路径，也就是 GO 语言的安装路径。
2. GOPATH：若干工作区目录的路径。是我们自己定义的工作空间。
3. GOBIN：GO 程序生成的可执行文件（executable file）的路径。

**设置 GOPATH 的意义是什么？**


### go get 安装包失败怎么办

错误信息的表现

```shell script
go get github.com/stretchr/testify: module github.com/stretchr/testify: Get "https://proxy.golang.org/github.com/stretchr/testify/@v/list": dial tcp 216.58.200.241:443: connectex: A
connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
```

使用国内的安装 [proxy](https://github.com/goproxy/goproxy.cn)

#### 使用方法

限制条件: Go 1.13 and above (RECOMMENDED)

```shell script
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
