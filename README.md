# gin-demo

## 前言
> 本项目参考自 [煎鱼](https://github.com/eddycjy) 系列文章 [《Gin搭建Blog API's》](https://eddycjy.com/posts/go/gin/2018-02-11-api-01)
> 我重构了整体的目录结构和代码，加入自己的一些想法和实践，来练习相关知识点，增加自己的对知识点的掌控。


## 目录结构

```shell
.
├── README.md
├── api.http
├── conf
│   ├── app.ini
│   └── config.yml
├── docs                          # 接口文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── gen
├── go.mod
├── go.sum
├── img.png
├── main.go
├── middleware
│   ├── rabc.go
│   └── timter.go
├── models                          # 模型层 （数据库）
│   ├── article.go
│   ├── models.go
│   ├── rbac.go
│   └── tag.go
├── pkg
│   ├── e
│   ├── errors                # 自定义错误包
│   ├── http                  # 自定义请求包
│   ├── rbac                  # casbin 权限校验
│   ├── setting               # 一些设置
│   ├── tools                 # 工具包
│   └── util
├── routers                         # 路由层
│   ├── api
│   └── router.go
├── run-go-vet.sh
├── runtime
├── service                         # 业务逻辑层
│   ├── service.go
│   └── service_tag.go
├── tests
│   ├── concatStr_test.go
│   ├── rabc_test.go
│   └── validator_test.go
├── tmp                             # 热更新临时文件
│   └── main.exe
├── types                           # 请求和响应返回结构体
    ├── request
    └── response

```

## 接口API文档（swagger）
![img.png](docs/img.png)


## 开发 TODO list

- [x] 目录结构重构
- [ ] 代码补充 & 重构
- [x] 接口文档 swagger 引入
- [ ] rbac 权限校验
- [ ] gRPC 调用 实践
- [ ] gRPC-gateway 实践
