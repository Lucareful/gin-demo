package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luenci/go-gin-example/config"
	db "github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/routers"
)

// 生成 swagger
//go:generate swag init -g routers/router.go --parseDependency --parseInternal --generatedTime --parseDepth 10
// 生成错误码
//go:generate codegen -type=int pkg/e/apiserver.go

func main() {
	// 初始化配置
	config.InitConf()
	conf := config.GetConf()

	// 初始化 Database
	db.InitDB()

	srv := &http.Server{
		Addr:           conf.Server.BindAddress,
		Handler:        routers.InitRouter(),
		ReadTimeout:    conf.Server.ReadTimeout * time.Second,
		WriteTimeout:   conf.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		log.Printf("server is runing port: %d\n", conf.Server.BindAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGSEGV,
		syscall.SIGABRT,
		syscall.SIGILL,
		syscall.SIGFPE,
		os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// Addr：监听的 TCP 地址，格式为:8000
// Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
// TLSConfig：安全传输层协议（TLS）的配置
// ReadTimeout：允许读取的最大时间
// ReadHeaderTimeout：允许读取请求头的最大时间
// WriteTimeout：允许写入的最大时间
// IdleTimeout：等待的最大时间
// MaxHeaderBytes：请求头的最大字节数
// ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
// ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出）
