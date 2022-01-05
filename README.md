# gin-demo

## 前言
> 本项目参考自 [煎鱼](https://github.com/eddycjy) 系列文章 [《Gin搭建Blog API's》](https://eddycjy.com/posts/go/gin/2018-02-11-api-01)
> 我重构了整体的目录结构和代码，加入自己的一些想法和实践，来练习相关知识点，增加自己的对知识点的掌控。


## 目录结构

```shell
.
├── CHANGELOG.md
├── README.md
├── api.http
├── config
│   ├── app.go
│   ├── app.ini
│   ├── app.yaml
│   └── app_test.go
├── docs
│   ├── docs.go
│   ├── img.png
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── cors.go
│   ├── jwt.go
│   ├── rabc.go
│   └── timter.go
├── models
│   ├── article.go
│   ├── jwt.go
│   ├── models.go
│   ├── rbac.go
│   ├── tag.go
│   └── user.go
├── pkg
│   ├── e
│   ├── http
│   ├── joint
│   ├── rbac
│   └── tools
├── routers
│   ├── api
│   └── router.go
├── runtime
├── service
│   ├── service.go
│   ├── serviceArticle.go
│   ├── serviceTag.go
│   └── serviceUser.go
├── tests
│   ├── concatStr_test.go
│   ├── jwt_test.go
│   ├── rabc_test.go
│   └── validator_test.go
├── types
├── request
└── response



```

## 接口API文档（swagger）
![img.png](docs/img.png)


## 开发 TODO list

- [x] 目录结构重构

- [ ] 代码补充 & 重构(80%)
- [x] [pkg 模块抽离](https://github.com/luenci)
- [x] [oauth2 认证服务](https://github.com/Lucareful/oauth2-server)
- [x] JWT token 签名认证
- [x] 接口文档 swagger 引入
- [ ] websocket 流式日志 & eventStream
- [ ] rbac 权限校验
- [x] [gRPC 实践](https://github.com/Lucareful/grpc-demo)
- [x] [gRPC-gateway 实践](https://github.com/Lucareful/grpc-getway-demo)
