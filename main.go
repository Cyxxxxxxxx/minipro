package main

import (
	"context"
	"fmt"
	"log"
	"minipro/controller"
	"minipro/logger"
	"minipro/mysql"
	"minipro/routes"
	"minipro/service"
	"minipro/setting"
	"minipro/tool"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title 小程序接口文档
// @version 1.0
// @description 宿舍管理小程序

// @contact.name  CCC
// @contact.email 2642427375@qq.com

// @host 127.0.0.1
// @BasePath :8080
func main() {
	//127.0.0.1
	//1、配置信息初始化
	err := setting.Init()
	if err != nil {
		fmt.Printf("config load failed err :", err)
	}

	//2、初始化日志
	err = logger.Init(setting.Conf.LogConfig, setting.Conf.Mode)
	if err != nil {
		fmt.Printf("logger load failed err :", err)
	}
	defer zap.L().Sync() // 延迟注册让缓冲区的数据也写入日志

	// 初始化时区为cst
	tool.InitInMain()

	//3、初始化数据库连接
	err = mysql.Init(setting.Conf.MySQLConfig)
	if err != nil {
		zap.L().Error("database connect  failed err :", zap.Error(err))
	} else {
		zap.L().Info("database01 connect success !")
	}
	defer mysql.Close()

	// 开启定时任务后台(消息推送)
	go controller.DataPushTaskHandle()

	//注册gin框架内置的校验器使用的翻译器
	if err = service.InitTrans("zh"); err != nil {
		zap.L().Error("init trans failed, err:%v\n", zap.Error(err))
		return
	}

	//5、注册路由
	r := routes.Setup(setting.Conf.Mode)
	r.Run()

	//6、启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err = srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	log.Println("Server exiting")
}
