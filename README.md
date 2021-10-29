# goweb-gin-demo
go web脚手架

# [web框架gin](https://gin-gonic.com/zh-cn/docs/introduction/)
## 特性
#### 快速
基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。

#### 支持中间件
传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，GZIP，最终操作 DB。

#### Crash 处理
Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！

#### JSON 验证
Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。

#### 路由组
更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。

#### 错误管理
Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。

#### 内置渲染
Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。

#### 可扩展性
新建一个中间件非常简单，去查看 [示例代码](https://gin-gonic.com/zh-cn/docs/examples/) 吧。

## 服务创建及启动

[官方demo](https://gin-gonic.com/zh-cn/docs/quickstart/)  
把ping映射在参数为(c *gin.Context)的方法  
```
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

endless 创建服务  
```
func main() {
	mux1 := mux.NewRouter()
	mux1.HandleFunc("/hello", handler).
		Methods("GET")

	srv := endless.NewServer("localhost:4244", mux1)
	srv.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1] = append(
		srv.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1],
		preSigUsr1)
	srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1] = append(
		srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1],
		postSigUsr1)

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 4244 stopped")

	os.Exit(0)
}
```


# 通过Swagger测试接口
[官方地址](https://github.com/swaggo) ,其中包含`swag`可执行程序和`gin-swagger`go web模块, [使用手册](https://github.com/swaggo/gin-swagger)  

```
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context)  {
	g.JSON(http.StatusOK,"helloworld")
}
```  

生成文档`swag init` 
```
docs/
├── docs.go
├── swagger.json
└── swagger.yaml
```

> 需要注意把docs文件夹导入到代码中 `import "github.com/ymm135/goweb-gin-demo/docs"`  

# 验证码获取及校验  
首选通过`/base/captcha`获取验证码,其中包含验证码图片地址:`data:image/png;...`  
```
{
    "code":0,
    "data":{
        "captchaId":"hrzSvHSdGo4Emm9oG9OY",
        "picPath":"data:image/png;base64,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"
    },
    "msg":"验证码获取成功"
}
```

在用户登录时，请求的数据为:  
```
{
	  "username": "admin",
	  "password": "123456",
	  "captcha": "153842",
	  "captchaId": "hrzSvHSdGo4Emm9oG9OY"
}
```  
携带着验证码及验证码的Id, 框架通过Id去匹配输入的验证码。  
登录城后后返回信息:  
```
{
    "code":0,
    "data":{
        "user":{
            "ID":1,
            "CreatedAt":"2021-10-25T14:53:31Z",
            "UpdatedAt":"2021-10-25T14:53:31Z",
            "uuid":"8e600b7f-3297-4979-a445-d218205ef9a6",
            "userName":"admin",
            "nickName":"超级管理员",
            "sideMode":"dark",
            "headerImg":"https:///qmplusimg.henrongyi.top/gva_header.jpg",
            "baseColor":"#fff",
            "activeColor":"#1890ff",
            "authorityId":"888",
            "authority":{
                "CreatedAt":"2021-10-25T14:53:31Z",
                "UpdatedAt":"2021-10-25T14:53:31Z",
                "DeletedAt":null,
                "authorityId":"888",
                "authorityName":"普通用户",
                "parentId":"0",
                "dataAuthorityId":null,
                "children":null,
                "defaultRouter":"dashboard"
            },
            "authorities":[
                {
                    "CreatedAt":"2021-10-25T14:53:31Z",
                    "UpdatedAt":"2021-10-25T14:53:31Z",
                    "DeletedAt":null,
                    "authorityId":"888",
                    "authorityName":"普通用户",
                    "parentId":"0",
                    "dataAuthorityId":null,
                    "children":null,
                    "defaultRouter":"dashboard"
                },
                {
                    "CreatedAt":"2021-10-25T14:53:31Z",
                    "UpdatedAt":"2021-10-25T14:53:31Z",
                    "DeletedAt":null,
                    "authorityId":"8881",
                    "authorityName":"普通用户子角色",
                    "parentId":"888",
                    "dataAuthorityId":null,
                    "children":null,
                    "defaultRouter":"dashboard"
                },
                {
                    "CreatedAt":"2021-10-25T14:53:31Z",
                    "UpdatedAt":"2021-10-25T14:53:31Z",
                    "DeletedAt":null,
                    "authorityId":"9528",
                    "authorityName":"测试角色",
                    "parentId":"0",
                    "dataAuthorityId":null,
                    "children":null,
                    "defaultRouter":"dashboard"
                }
            ]
        },
        "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOGU2MDBiN2YtMzI5Ny00OTc5LWE0NDUtZDIxODIwNWVmOWE2IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTYzNjAyMDY2MSwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjM1NDE0ODYxfQ.-E9YNI6YV3XLbC5GgXiftTBI7A5ZiIheH5dqwgwnaDA",
        "expiresAt":1636020661000
    },
    "msg":"登录成功"
}
```


# 获取菜单项
api接口`/api/menu/getMenu`  

头文件携带信息:
```
x-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOGU2MDBiN2YtMzI5Ny00OTc5LWE0NDUtZDIxODIwNWVmOWE2IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTYzNjAyMDY2MSwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjM1NDE0ODYxfQ.-E9YNI6YV3XLbC5GgXiftTBI7A5ZiIheH5dqwgwnaDA
x-user-id: 1
```  



## 使用[Casbin](https://github.com/casbin/casbin) 控制权限
Casbin is a powerful and efficient open-source access control library. It provides support for enforcing authorization based on various access control models.  

可以 [在线](https://casbin.org/editor/) 书写规则:    

![在线写casbin规则](./resource/md_res/casbin.png)  

创建一个Casbin enforcer  
```
import "github.com/casbin/casbin/v2"

e, err := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
``` 

检查权限  
```
sub := "alice" // the user that wants to access a resource.
obj := "data1" // the resource that is going to be accessed.
act := "read" // the operation that the user performs on the resource.

ok, err := e.Enforce(sub, obj, act)

if err != nil {
    // handle err
}

if ok == true {
    // permit alice to read data1
} else {
    // deny the request, show an error
}

// You could use BatchEnforce() to enforce some requests in batches.
// This method returns a bool slice, and this slice's index corresponds to the row index of the two-dimensional array.
// e.g. results[0] is the result of {"alice", "data1", "read"}
results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
```





# 文件上传及下载  

# 反射reflect 
[官方文档](https://pkg.go.dev/reflect)  

包反射实现运行时反射，允许程序操作任意类型的对象。典型的用法是取一个静态类型 interface{} 的值，通过调用 TypeOf 提取其动态类型信息，返回一个 Type。  

对 ValueOf 的调用返回一个表示运行时数据的值。Zero 接受一个 Type 并返回一个表示该类型的零值的 Value。    

有关 Go 中反射的介绍，请参阅 [“反射定律”](https://golang.org/doc/articles/laws_of_reflection.html)   

动态调用有参数的函数  
```
type T struct{}

func main() {
    name := "Do"
    t := &T{}
    a := reflect.ValueOf(1111)
    b := reflect.ValueOf("world")
    in := []reflect.Value{a, b}
    reflect.ValueOf(t).MethodByName(name).Call(in)
}

func (t *T) Do(a int, b string) {
    fmt.Println("hello" + b, a)
}
``` 

# 疑问及拓展
#### 为什么函数或方法中变量名很多都是大写字母开始的?
- 任何需要对外暴露的名字必须以大写字母开头，不需要对外暴露的则应该以小写字母开头。     
但是gin-vue-admin代码中有很多大小，连方法参数都是大写字母开始的?  
```
# 函数中的变量Router
Router := initialize.Routers()

# 方法参数Router
func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
```

#### unsupported Scan, storing driver.Value type []uint8 into type *time.Time
需要在连接数据库时增加参数: 
```
db, err := sqlx.Connect("mysql", "myuser:mypass@tcp(127.0.0.1:3306)/mydb?parseTime=true") 
```

Sets the location for time.Time values (when using parseTime=true). "Local" sets the system's location. See [time.LoadLocation](https://pkg.go.dev/time#LoadLocation) for details.  

  
####  Error 1075: Incorrect table definition; there can be only one auto column and it must be defined as a key

日志信息
> 2021/10/29 11:44:55 /Users/zero/go/pkg/mod/github.com/casbin/gorm-adapter/v3@v3.4.4/adapter.go:373 
> Error 1075: Incorrect table definition; there can be only one auto column and it must be defined as a key
>  [7.259ms] [rows:0] ALTER TABLE `casbin_rule` ADD `id` bigint unsigned AUTO_INCREMENT  

断点调试gorm-adapter/v3@v3.4.4/adapter.go, 最后上网搜索发现版本问题。 把版本都降低了
	github.com/casbin/casbin/v2 v2.11.0
	github.com/casbin/gorm-adapter/v3 v3.0.2
	

  













  





