# 目录
- [goweb-gin-demo](#goweb-gin-demo)
- [web框架gin](#web框架gin)
    - [特性](#特性)
	- [服务创建及启动](#服务创建及启动)
- [gorm](#gorm)
	- [概述](#概述)
	- [模型定义](#模型定义)
	- [连接数据库](#连接数据库)
		- [连接池](#连接池)
	- [CRUD](#crud)
		- [基本操作](#基本操作)
		- [创建钩子](#创建钩子)
- [通过Swagger测试接口, 中文文档](#通过swagger测试接口-中文文档)
- [验证码获取及校验](#验证码获取及校验)
- [获取菜单项](#获取菜单项)
- [使用casbin 控制权限](#使用casbin-控制权限)
- [文件上传及下载](#文件上传及下载)
- [代码自动生成](#代码自动生成)
- [反射reflect](#反射reflect)
- [GoWeb 发布](#goweb-发布)
- [疑问及拓展](#疑问及拓展)
    - [为什么函数或方法中变量名很多都是大写字母开始的?](#为什么函数或方法中变量名很多都是大写字母开始的)
    - [unsupported Scan, storing driver.Value type []uint8 into type *time.Time](#unsupported-scan-storing-drivervalue-type-uint8-into-type-timetime)
    - [Error 1075: Incorrect table definition; there can be only one auto column and it must be defined as a key](#error-1075-incorrect-table-definition-there-can-be-only-one-auto-column-and-it-must-be-defined-as-a-key)
    - [查看http请求详情](#查看http请求详情)
  
# goweb-gin-demo
go web脚手架, [数据库及表结构](./resource/sql/weekly_report.sql)  

# [web框架gin](https://gin-gonic.com/zh-cn/docs/introduction/) 
[中文参考文档](https://www.kancloud.cn/shuangdeyu/gin_book/949411) 
## 特性
- 快速
基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。

- 支持中间件
传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，GZIP，最终操作 DB。

- Crash 处理
Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！

- JSON 验证
Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。

- 路由组
更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。

- 错误管理
Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。

- 内置渲染
Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。

- 可扩展性
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

# [gorm](https://gorm.io/zh_CN/docs/index.html)   
## 概述
- 全功能 ORM
- 关联 (Has One，Has Many，Belongs To，Many To Many，多态，单表继承)属于、多对多、多态、单表继承）
- Create，Save，Update，Delete，Find 中钩子方法保存/更新/删除/查找）
- 支持 Preload、Joins 的预加载，Joins
- 事务，嵌套事务，Save Point，Rollback To Saved Point、回滚到保存点
- Context、预编译模式、DryRun 模式DryRun 模式
- 批量插入，FindInBatches，Find/Create with Map，使用 SQL 表达式、Context Valuer 进行 CRUDBatches、使用 Map 查找/创建、使用 SQL Expr 和 Context Valuer 进行 CRUD
- SQL 构建器，Upsert，数据库锁，Optimizer/Index/Comment Hint，命名参数，子查询、Upsert、锁定、优化器/索引/注释提示、命名参数、子查询
- 复合主键，索引，约束
- Auto Migration
- 自定义 Logger
- 灵活的可扩展插件 API：Database Resolver（多数据库，读写分离）、Prometheus…PI：数据库解析器（多个数据库，读/写拆分）/ Prometheus...
- 每个特性都经过了测试的重重考验
- 开发者友好
  
## 模型定义 
```
type User struct {
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
```

GORM 使用结构体名的 蛇形命名 作为表名。对于结构体 User，根据约定，其表名为 users  
您可以实现 Tabler 接口来更改默认表名，例如：  
```
type Tabler interface {
    TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
  return "profiles"
}
```

您可以使用 Table 方法临时指定表名，例如：
```
// 根据 User 的字段创建 `deleted_users` 表
db.Table("deleted_users").AutoMigrate(&User{})

// 从另一张表查询数据
var deletedUsers []User
db.Table("deleted_users").Find(&deletedUsers)
// SELECT * FROM deleted_users;

db.Table("deleted_users").Where("name = ?", "jinzhu").Delete(&User{})
// DELETE FROM deleted_users WHERE name = 'jinzhu';
```

## 连接数据库  
```
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

### 连接池
GORM 使用 database/sql 维护连接池
```
sqlDB, err := db.DB()

// SetMaxIdleConns 设置空闲连接池中连接的最大数量
sqlDB.SetMaxIdleConns(10)

// SetMaxOpenConns 设置打开数据库连接的最大数量。
sqlDB.SetMaxOpenConns(100)

// SetConnMaxLifetime 设置了连接可复用的最大时间。
sqlDB.SetConnMaxLifetime(time.Hour)
```

## CRUD 
### 基本操作
```
#  创建记录
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // 通过数据的指针来创建

user.ID             // 返回插入数据的主键
result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数
----------------------------------------------------------------------

#  查询
// 获取第一条记录（主键升序）
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取一条记录，没有指定排序字段
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// 获取最后一条记录（主键降序）
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;
----------------------------------------------------------------------

#  更新
user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111
----------------------------------------------------------------------

# 删除
// Email 的 ID 是 `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// 带额外条件的删除
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";
----------------------------------------------------------------------
```
### 创建钩子  
GORM 允许用户定义的钩子有 BeforeSave, BeforeCreate, AfterSave, AfterCreate 创建记录时将调用这些钩子方法，请参考 Hooks 中关于生命周期的详细信息
```
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}
```


# 通过Swagger测试接口, [中文文档](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)  
[官方地址](https://github.com/swaggo) ,其中包含`swag`可执行程序和`gin-swagger`go web模块, [使用手册](https://github.com/swaggo/gin-swagger)  
查看地址为: **http://IP:8888/swagger/index.html**  

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

> 需要注意把docs文件夹导入到代码中 `import "goweb-gin-demo/docs"`  

- **如何获取到方法的注释呢?**  


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

jwt token  
<br>
<div align=center>
<img src="./resource/md_res/jwt.png" width="60%" height="60%" title="JWT Token"></img>
</div>

<br>
**用户-角色权限-菜单项|API|资源对应关系图**   
<br>
   
![用户授权](./resource/md_res/user-auth.jpg)  


# 获取菜单项
api接口`/api/menu/getMenu`  

头文件携带信息:
```
x-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiOGU2MDBiN2YtMzI5Ny00OTc5LWE0NDUtZDIxODIwNWVmOWE2IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTYzNjAyMDY2MSwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjM1NDE0ODYxfQ.-E9YNI6YV3XLbC5GgXiftTBI7A5ZiIheH5dqwgwnaDA
x-user-id: 1
```  

- **角色对应着菜单项和API接口**   

通过userid获取权限id(`SELECT * FROM `sys_user_authority` WHERE `sys_user_authority`.`sys_user_id` = 1`)      
```
+-------------+----------------------------+
| sys_user_id | sys_authority_authority_id |
+-------------+----------------------------+
|           1 | 888                        |   普通用户
|           1 | 8881                       |   普通子用户
|           1 | 9528                       |   测试用户
+-------------+----------------------------+
```   


通过权限ID获取权限(`SELECT * FROM `sys_authorities` WHERE `sys_authorities`.`authority_id` IN ('888','8881','9528')`)  
```
+---------------------+---------------------+------------+--------------+----------------+-----------+----------------+
| created_at          | updated_at          | deleted_at | authority_id | authority_name | parent_id | default_router |
+---------------------+---------------------+------------+--------------+----------------+-----------+----------------+
| 2021-10-25 14:53:31 | 2021-10-25 14:53:31 | NULL       | 888          | ????           | 0         | dashboard      |
| 2021-10-25 14:53:31 | 2021-10-25 14:53:31 | NULL       | 8881         | ???????        | 888       | dashboard      |
| 2021-10-25 14:53:31 | 2021-10-25 14:53:31 | NULL       | 9528         | ????           | 0         | dashboard      |
+---------------------+---------------------+------------+--------------+----------------+-----------+----------------+
```             

返回数据样本
```
{
    "code":0,
    "data":{
        "menus":[
            {
                "ID":1,
                "CreatedAt":"2021-11-01T11:14:12Z",
                "UpdatedAt":"2021-11-01T11:14:14Z",
                "parentId":"0",
                "path":"",
                "name":"周报",
                "hidden":false,
                "component":"",
                "sort":1,
                "meta":{
                    "keepAlive":false,
                    "defaultMenu":false,
                    "title":"周报",
                    "icon":"",
                    "closeTab":false
                },
                "authoritys":null,
                "menuId":"1",
                "children":[
                    {
                        "ID":2,
                        "CreatedAt":"2021-11-01T11:16:40Z",
                        "UpdatedAt":"2021-11-01T11:16:43Z",
                        "parentId":"1",
                        "path":"",
                        "name":"查看周报",
                        "hidden":false,
                        "component":"",
                        "sort":1,
                        "meta":{
                            "keepAlive":false,
                            "defaultMenu":false,
                            "title":"查看周报",
                            "icon":"",
                            "closeTab":false
                        },
                        "authoritys":null,
                        "menuId":"2",
                        "children":null,
                        "parameters":[

                        ]
                    },
                    {
                        "ID":3,
                        "CreatedAt":"2021-11-01T11:16:40Z",
                        "UpdatedAt":"2021-11-01T11:16:43Z",
                        "parentId":"1",
                        "path":"",
                        "name":"写周报",
                        "hidden":false,
                        "component":"",
                        "sort":2,
                        "meta":{
                            "keepAlive":false,
                            "defaultMenu":false,
                            "title":"写周报",
                            "icon":"",
                            "closeTab":false
                        },
                        "authoritys":null,
                        "menuId":"3",
                        "children":null,
                        "parameters":[

                        ]
                    },
                    {
                        "ID":4,
                        "CreatedAt":"2021-11-01T11:16:40Z",
                        "UpdatedAt":"2021-11-01T11:16:43Z",
                        "parentId":"1",
                        "path":"",
                        "name":"统计导出",
                        "hidden":false,
                        "component":"",
                        "sort":3,
                        "meta":{
                            "keepAlive":false,
                            "defaultMenu":false,
                            "title":"统计导出",
                            "icon":"",
                            "closeTab":false
                        },
                        "authoritys":null,
                        "menuId":"4",
                        "children":null,
                        "parameters":[

                        ]
                    },
                    {
                        "ID":5,
                        "CreatedAt":"2021-11-01T11:16:40Z",
                        "UpdatedAt":"2021-11-01T11:16:43Z",
                        "parentId":"1",
                        "path":"",
                        "name":"模板编辑",
                        "hidden":false,
                        "component":"",
                        "sort":4,
                        "meta":{
                            "keepAlive":false,
                            "defaultMenu":false,
                            "title":"模板编辑",
                            "icon":"",
                            "closeTab":false
                        },
                        "authoritys":null,
                        "menuId":"5",
                        "children":null,
                        "parameters":[

                        ]
                    }
                ],
                "parameters":[

                ]
            },
            {
                "ID":6,
                "CreatedAt":"2021-11-01T11:16:40Z",
                "UpdatedAt":"2021-11-01T11:16:43Z",
                "parentId":"0",
                "path":"",
                "name":"设置",
                "hidden":false,
                "component":"",
                "sort":2,
                "meta":{
                    "keepAlive":false,
                    "defaultMenu":false,
                    "title":"设置",
                    "icon":"",
                    "closeTab":false
                },
                "authoritys":null,
                "menuId":"6",
                "children":[
                    {
                        "ID":7,
                        "CreatedAt":"2021-11-01T11:16:40Z",
                        "UpdatedAt":"2021-11-01T11:16:43Z",
                        "parentId":"6",
                        "path":"",
                        "name":"用户密码",
                        "hidden":false,
                        "component":"",
                        "sort":1,
                        "meta":{
                            "keepAlive":false,
                            "defaultMenu":false,
                            "title":"用户密码",
                            "icon":"",
                            "closeTab":false
                        },
                        "authoritys":null,
                        "menuId":"7",
                        "children":null,
                        "parameters":[

                        ]
                    },
                    {
                        "ID":8,
                        "CreatedAt":"2021-11-01T11:16:40Z",
                        "UpdatedAt":"2021-11-01T11:16:43Z",
                        "parentId":"6",
                        "path":"",
                        "name":"统计规则",
                        "hidden":false,
                        "component":"",
                        "sort":2,
                        "meta":{
                            "keepAlive":false,
                            "defaultMenu":false,
                            "title":"统计规则",
                            "icon":"",
                            "closeTab":false
                        },
                        "authoritys":null,
                        "menuId":"8",
                        "children":null,
                        "parameters":[

                        ]
                    },
                    {
                        "ID":9,
                        "CreatedAt":"2021-11-01T11:16:40Z",
                        "UpdatedAt":"2021-11-01T11:16:43Z",
                        "parentId":"6",
                        "path":"",
                        "name":"用户管理",
                        "hidden":false,
                        "component":"",
                        "sort":3,
                        "meta":{
                            "keepAlive":false,
                            "defaultMenu":false,
                            "title":"用户管理",
                            "icon":"",
                            "closeTab":false
                        },
                        "authoritys":null,
                        "menuId":"9",
                        "children":null,
                        "parameters":[

                        ]
                    }
                ],
                "parameters":[

                ]
            }
        ]
    },
    "msg":"获取成功"
}
```



# 使用[casbin](https://github.com/casbin/casbin) 控制权限
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

文件上传及下载code
```
func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的路径
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}
```

```
// @Tags FileUploadAndDownload
// @Summary 文件下载
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param fileName query string true "待下载的文件名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"文件下载成功"}"
// @Router /fileUploadAndDownload/download [get]
func (u *FileUploadAndDownloadApi) DownloadFile(c *gin.Context) {
	fileName := c.Query("fileName")

	//待下载的文件是存储本地的,如果存储在云盘，直接访问链接即可
	//如果没有记录，也会报错，record not found
	err, fileInfo := fileUploadAndDownloadService.FindFile(fileName)

	if err != nil {
		global.GLOBAL_LOG.Error("文件未找到!", zap.Any("err", err))
		c.Redirect(http.StatusFound, "/404")
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileInfo.Name)
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(fileInfo.Url)
	return
}
```

文件上传成功返回:  
```
{
    "code": 0,
    "data": {
        "file": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "name": "demo.c",
            "url": "/Users/zero/Documents/uploads/file/fe01ce2a7fbac8fafaed7c982a04e229_20211102170427.c",
            "tag": "c",
            "key": "fe01ce2a7fbac8fafaed7c982a04e229_20211102170427.c"
        }
    },
    "msg": "上传成功"
}
```

数据库信息:  
```
mysql> select * from file_upload_and_downloads;
+----+---------------------+---------------------+------------+--------+--------------------------------------------------------------------------------------+------+---------------------------------------------------+
| id | created_at          | updated_at          | deleted_at | name   | url                                                                                  | tag  | key                                               |
+----+---------------------+---------------------+------------+--------+--------------------------------------------------------------------------------------+------+---------------------------------------------------+
|  3 | 2021-11-02 09:04:27 | 2021-11-02 09:04:27 | NULL       | demo.c | /Users/zero/Documents/uploads/file/fe01ce2a7fbac8fafaed7c982a04e229_20211102170427.c | c    | fe01ce2a7fbac8fafaed7c982a04e229_20211102170427.c |
+----+---------------------+---------------------+------------+--------+--------------------------------------------------------------------------------------+------+---------------------------------------------------+
```

# 代码自动生成
gin-vue-admin提供代码自动生成功能，选择数据库表结构，可直接生成前端代码及后端代码，很方便  
比如周报的表结构如下:  
```
{
    "header":"周报",
    "contents":[
        {"title":"", "content":""},
        {"title":"", "content":""}
    ],
    "attachments":[
         {"pic":"fe01ce2a7fbac8fafaed7c982a04e229_20211102170427.png"},
         {"file":"2db0df442ea6d92d75657712a29e5604_20211102201017.txt"}
    ]
}
```

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

# GoWeb 发布
go build编译出moduleName的可执行文件，加载配置文件config.yaml即可运行。    


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
	

#### 查看http请求详情

通过查看gin.Context>>Request>>(Body io.ReadCloser)>>src>>R>>buf=>点击view就能看到http详情。  

<br>
<div align=center>
<img src="./resource/md_res/gin-http-res.jpg" width="70%" height="70%" title="Go Http请求"></img>
</div>
<br>  



	

	

  













  





