<a name="unreleased"></a>
## [Unreleased]


<a name="v0.5"></a>
## v0.5 - 2021-11-11
### Add
- 分页器开发
- 补充 tag 逻辑
- 适配 linux 平台下 air 配置更新 & article 接口逻辑增加
- generate swagger & error code
- Tag service 逻辑重构 & 完善
- Get 接口逻辑
- WithReponse 定义
- 引入三方错误码 WithStack
- air 热更新 & casbin table
- 基于 casbin 的 rabc 权限校验
- ConcatStr & 单元测试
- code error 自定义
- code error 自定义
- gin ShouldBindJSON 应用
- article 接口逻辑实现
- tag 表设置随时间更新字段
- validator 校验 & 增加tags
- gormt 和 validator 实践
- 路由注册
- 增加 timter 请求时间统计中间件
- 目录结构完善

### Delete
- 删除冗余文件

### Feat
- 新增swag自动生成 swagger 文档&删除api-test接口
- 增加自定义WithResponse

### Init
- 初始化提交

### Refactor
- 重构response
- 重构 tag 参数校验
- 拆分 service & 接口注册
- 拆分 service 逻辑层
- 重构 Tag list 接口,拆分 request 和 response

### Refator
- 生成错误码
- 自定义错误返回

### Style
- pre-commit 检查

### Update
- 更新代码结构和补全tag逻辑
- 优雅退出或重启功能
- 更新README
- 更新表结构
- 修复字符串拼接产生的内存复制
- 更新README
- 更新错误码
- 更新 README
- 更新 README
- 抽离路由


[Unreleased]: https://github.com/Lucareful/gin-demo/compare/v0.5...HEAD
