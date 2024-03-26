package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Astrotrain-ali/cmdb/apps"
	"github.com/Astrotrain-ali/cmdb/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
)

var (
	confType string
	confFile string
	confETCD string
)

// 程序的启动时 组装都在这里进行
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端api",
	Long:  "启动 demo 后端api",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}

		// 初始化全局日志logger
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 加载Host Service的实体类
		// host service的具体实现
		//service := impl.NewHostServiceImpl()

		// 注册HostService的实例到IOC
		// 采用：_ "gitee.com/max-astrotrain/restful-api-demo/apps/host/impl" 完成注册
		//apps.HostService = impl.NewHostServiceImpl()

		// 如何执行HostService的config方法
		// 因为apps.HostService是host.Service接口，并没有保存实例初始化（Config）的方法
		apps.InitImpl()

		// // 提供一个Gin Router对外提供服务
		// g := gin.Default()
		// // 注册IOC的所有http handler
		// apps.InitGin(g)
		// g.Run(conf.C().App.HttpAddr())
		svc := newManager()

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go svc.WaitStop(ch)
		return svc.Start()
	},
}

func newManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    zap.L().Named("cli"),
	}
}

// 用于管理所有需要启动的服务
// 1. HTTP服务的启动
type manager struct {
	http *protocol.HttpService
	l    logger.Logger
}

func (m *manager) Start() error {
	return m.http.Start()
}

// 处理来自外部的中断信号，比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.l.Infof("received signal: %s", v)
			m.http.Stop()
		}
	}
}

//	初始化logger实例
//
// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 根据Config里面的日志配置，来配置全局logger对象
	lc := conf.C().Log

	// 解析日志level配置
	// DebugLevel: "debug",
	// InfoLevel:  "info",
	// WarnLevel:  "warning",
	// ErrorLevel: "error",
	// FatalLevel: "fatal",
	// PanicLevel: "panic",
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	// 使用默认配置初始化Logger的全局配置
	zapConfig := zap.DefaultConfig()

	// 配置日志的level级别
	zapConfig.Level = level

	// 程序每启动一次，不必都生成一个新日志文件
	zapConfig.Files.RotateOnStartup = false

	// 配置日志的输出方式
	switch lc.To {
	case conf.ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没有把日志输入输出到文件
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	// 配置日志的输出格式：
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}

	// 把配置应用到全局logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
