# yan_api

这是一个开发项目的分享视频。做一个社交交友项目。把其中的全部流程分享出来。
yan_api: 言项目的后端api程序

![项目思维导图](https://github.com/royromny/yan-api/blob/master/assets/img/%E9%A1%B9%E7%9B%AE%E6%80%9D%E7%BB%B4%E5%AF%BC%E5%9B%BE.jpg?raw=true)

## 言 项目的资料
### 语雀资料包括：
[言项目的资料](https://www.yuque.com/books/share/e96b3f14-1746-4780-b4b1-c09fa23820b7)

[项目思维导图](https://www.yuque.com/docs/share/2b643335-2fad-4582-8e6b-249bf01239d6)
[视频流程目录](https://www.yuque.com/docs/share/b8058c87-3eaa-442f-a0f0-e3765dde183b)

## 视频实况分享流程
### 已完成
- a. 【yan】#2a-后端语言： golang ，使用轻量级框架： gin ，使用落地代码框架： singo。[头条](https://www.ixigua.com/i6859617878241444366/) | [b站](https://www.bilibili.com/video/BV15Z4y1K7sC/)
- b. 【yan】#2b-调整封装框架的返回。[头条](https://www.ixigua.com/i6859648336010576398/) | [b站](https://www.bilibili.com/video/BV1uA411e7X9/)
- c. 【yan】#2c-项目配置文件，处理软件配置方法，用ymal。[头条](https://www.ixigua.com/i6859999281500127755/) | [b站](https://www.bilibili.com/video/BV1Vf4y197Zs/)
- d. 【yan】#2d-链接数据库，gorm链接数据库和自动迁移，mysql。[头条](https://www.ixigua.com/i6860010156923552270/) | [b站](https://www.bilibili.com/video/BV1xZ4y1M76D/)
- e. 【yan】#2e-日志管理，logger。[头条](https://www.ixigua.com/i6860435086181925387/) | [b站](https://www.bilibili.com/video/BV1Gp4y1v7TE/)
- f. 【yan】#2f-接收请求，GET、POST、PUT、DELETE，接收json数据并反序列化到结构体并验证。[头条](https://www.ixigua.com/i6860475245216662019/) | [b站](https://www.bilibili.com/video/BV15C4y1t7ZW/)
- g. 【yan】#2g-模型中基本的增删改查。[头条](https://www.ixigua.com/i6860806601721250311/) | [b站](https://www.bilibili.com/video/BV1Kt4y1Q7nA/)
- h. 【yan】#2h-golang编程，RESTful返回结果序列化json。

### 计划中
- i. 【yan】#2h-请求接口的监权，session使用，RBAC模型。

## 目的

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API

## 特色

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
6. 自行实现了国际化i18n的一些基本功能
7. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证

本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/user/register```用户注册接口
3. 实现了```/api/v1/user/login```用户登录接口
4. 实现了```/api/v1/user/me```用户资料接口(需要登录后获取session)
5. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件


## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init yan-api
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go mod tidy
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)