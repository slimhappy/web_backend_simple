# 从hello world到简单的CRUD后台框架
## 1. 构建一个简单的go工程
```shell
$ makedir web_backend_simple
// 初始化 git
$ git init
// 使用 gomod 管理依赖
$ go mod init github.com/slimhappy/web_backend_simple
// 新建 mian.go
$ touch main.go
```
## 2. Hello World 代码
```GO
package main
	
import "fmt"

func main() {
    fmt.Println("hello world")
}
```
## 3. 编译并运行 Hello World 代码
```shell
$ go build main.go
$ ./mian
hello world
```

## 4. 下载gin框架
```shell
$ go get -u github.com/gin-gonic/gin
go: downloading github.com/gin-gonic/gin v1.7.7
go: github.com/gin-gonic/gin upgrade => v1.7.7
...//省略一堆的依赖下载
```

## 5. 参考gin的官方文档新建一个简单返回的接口
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
编译并运行
```shell
$ go build main.go
$ ./mian
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```
这里居然有一个WARNING！
```
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
```
不慌不慌，我们打开这个参考页面，并且得到解决方法，在代码里加上：
```
	r.SetTrustedProxies(nil)
```

## 6. 修改工程结构：
```

```