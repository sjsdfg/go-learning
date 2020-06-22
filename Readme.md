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

### go 测试框架

go 测试框架目前看 [testify](https://github.com/stretchr/testify) 是主流框架

```go
package yours

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

  // assert for nil (good for errors)
  assert.Nil(t, object)

  // assert for not nil (good when you expect something)
  if assert.NotNil(t, object) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(t, "Something", object.Value)

  }
}
```

### Introduce local variable in Gogland ide

[原答案地址](https://stackoverflow.com/questions/44657806/introduce-local-variable-in-gogland-ide)

You need to use CTRL+ALT+V for Windows / Linux (or CMD+ALT+V on OS X) or invoke the Refactor | Extract | Variable and then select the function call from the list and the variables will be inserted for you.


