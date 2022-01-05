module github.com/luenci/go-gin-example

go 1.15

require (
	github.com/casbin/casbin/v2 v2.37.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-playground/validator/v10 v10.9.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/luenci/errors v1.0.1
	github.com/marmotedu/errors v1.0.2
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/novalagung/gubrak v1.0.0
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/spf13/viper v1.9.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.3.2
	github.com/swaggo/swag v1.7.3
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	github.com/wxnacy/wgo v1.0.4
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.7
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.22.2
)

// 后续每新增一个本地应用目录，你都需要主动去 go.mod 文件里新增一条 replace
replace (
	github.com/luenci/go-gin-example/config => ./config
	github.com/luenci/go-gin-example/docs => ./docs
	github.com/luenci/go-gin-example/middleware => ./middleware
	github.com/luenci/go-gin-example/models => ./models
	github.com/luenci/go-gin-example/pkg => ./pkg
	github.com/luenci/go-gin-example/pkg/setting => ./pkg/setting
	github.com/luenci/go-gin-example/routers => ./routers
	github.com/luenci/go-gin-example/routers/api => ./api
	github.com/luenci/go-gin-example/service => ./service
	github.com/luenci/go-gin-example/test => ./test
	github.com/luenci/go-gin-example/types => ./types
)
