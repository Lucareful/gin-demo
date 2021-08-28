module github.com/luenci/go-gin-example

go 1.15

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/sys v0.0.0-20210820121016-41cdb8703e55 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.47.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/gorm v1.21.13
)

// 后续每新增一个本地应用目录，你都需要主动去 go.mod 文件里新增一条 replace
replace (
	github.com/luenci/go-gin-example/middleware => ./middleware
	github.com/luenci/go-gin-example/models => ./models
	github.com/luenci/go-gin-example/pkg => ./pkg
	github.com/luenci/go-gin-example/pkg/setting => ./pkg/setting
	github.com/luenci/go-gin-example/routers => ./routers
	github.com/luenci/go-gin-example/service => ./service
	github.com/luenci/go-gin-example/types => ./types
)
