# go-vue-admin

go-vue-admin 是一套基于go语言开源的后台管理系统。功能参考[诺依网站](http://www.ruoyi.vip/) ，前后端分离。

## 简介

- 前端采用vue3、[Element Plus](https://element-plus.org/zh-CN/) 、[RuoYi-Vue3](https://gitee.com/y_project/RuoYi-Vue)
- 后端采用gofrome 框架、mysql、redis、Jwt
- 实现了一键生成前后端代码，高效开发。

##  内置功能

1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6. 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7. 参数管理：对系统动态配置常用参数。
8. 通知公告：系统通知公告信息发布维护。
9. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
10. 登录日志：系统登录日志记录查询包含登录异常。
11. 在线用户：当前系统中活跃用户状态监控。
12. 定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
13. 代码生成：前后端代码的生成（go，js，vue）支持CRUD下载 。
14. 系统接口：根据业务代码自动生成相关的api接口文档。
15. 服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
16. 缓存监控：对系统的缓存信息查询，命令统计等。
17. 在线构建器：拖动表单元素生成相应的HTML代码。
18. 连接池监视：监视当前系统数据库连接池状态，可进行分析SQL找出系统性能瓶颈。

##  系统环境

golang ：go1.18+ 

数据库：mysql5.7+

缓存：redis3.0+

Node ： 12+

## 项目地址

github:

https://github.com/guyan0319/go-vue-admin

码云（国内）:

https://gitee.com/jason0319/go-vue-admin

## 快速开始

1、clone项目源代码

`git clone https://github.com/guyan0319/go-vue-admin`

2、安装 yarn

```
npm install -g yarn 
```

3、新建数据库名（如：gvadmindb） 、导入目录 manifest/sql/gvadmindb.sql

根据实际环境修改 hack/config.yaml

```
    dao:
      - link:            "mysql:root:123456@tcp(127.0.0.1:3306)/gvadmindb"
        tables:          ""
        tablesEx:        ""
        removePrefix:    "gf_"
        descriptionTag:  true
        noModelComment:  true
        path: "./internal/app/system"
```

manifest/config/config.yaml 中数据库配置和redis配置信息

```
# Database.
database:
  logger:
    level:   "all"
    stdout:  true
    Path: "resource/log/sql"

  default:
    link:   "mysql:root:123456@tcp(127.0.0.1:3306)/gvadmindb?charset=utf8mb4&parseTime=true&loc=Local"
    debug:  true
    charset: "utf8mb4" #数据库编码
    dryRun: false #空跑
    maxIdle: 10 #连接池最大闲置的连接数
    maxOpen: 10 #连接池最大打开的连接数
    maxLifetime: "30s" #(单位秒)连接对象可重复使用的时间长度
# Redis 配置示例
redis:
  # 单实例配置
  default:
    address: 127.0.0.1:6379
    db: 1
#    pass:    "password" # 在此配置密码, 没有可去掉
    idleTimeout: "60s" #连接最大空闲时间，使用时间字符串例如30s/1m/1d
    maxConnLifetime: "90s" #连接最长存活时间，使用时间字符串例如30s/1m/1d
    waitTimeout: "60s" #等待连接池连接的超时时间，使用时间字符串例如30s/1m/1d
    dialTimeout: "30s" #TCP连接的超时时间，使用时间字符串例如30s/1m/1d
    readTimeout: "30s" #TCP的Read操作超时时间，使用时间字符串例如30s/1m/1d
    writeTimeout: "30s" #TCP的Write操作超时时间，使用时间字符串例如30s/1m/1d
    maxActive: 100

```

4、启动服务端

```
go run main.go
```

如果在开发环境，热更新可使用

```
gf run main.go
```

5、启动前端，打开RuoYi-Vue3目录

```
yarn dev 
```

6、浏览器打开

```
http://localhost/login?redirect=/index
```

登录测试账户信息

账户：admin

密码：admin123

## 感谢

- gf框架 <https://github.com/gogf/gf>

- RuoYi-Vue3 https://gitee.com/y_project/RuoYi-Vue

## 注意

开发者模式下，登录不验证验证码，可随意字符。

如生产环境去掉//gmode.SetProduct()前面注释