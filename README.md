# DogeGo Mini

使用DogeGo Mini开发分布式WebAPI: 用最简单的架构, 实现支持分布式的够用的框架, 服务大量的用户.

本框架是DogeGo的精简版本，保留了DogeGo的高性能，去掉了一些对于简单API项目无用的东西.

### 目的
本项目采用了一些Golang里面比较流行的我试用过的组件，让大家快速搭建支持分布式的Restful Web API.

### 特色
本项目整合了开发API的基础类库:
1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件

本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/ping```心跳检查接口

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. models文件夹负责存储数据库模型和数据库操作相关的代码
3. services负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把models得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. utils一些通用的小工具
7. conf放一些静态存放的配置文件
8. global放Redis全局计数器、排行榜之类的分布式共享资源
9. modules放需要全局单例使用的模块, 类似于ASP.NetCore的IOC的单例模式
10. routers放路由信息

### 赞助
![赞助](https://probe.aiursoft.com/Download/Open/%E5%9B%BE%E5%BA%8A/1566653941311.jpg)
