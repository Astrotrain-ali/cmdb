package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查（编译器提供的语法检查）
// 这样写会造成conf.C()并准备好，造成conf.C().MySQL.GetDB()该方法panic

// 把对象的注册和对象的初始化独立出来
var impl = &HostServiceImpl{}

// NewHostServiceImpl 保证调用该函数之前，全局conf对象已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service的子loggger
		// 封装的zap让其满足logger接口
		// 为什么要封装：
		// 		1. logger全局实例
		// 		2. logger level的动态调整，logrus不支持level动态调整
		// 		3. 加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

// 只需要保证 全局对象Config和全局Logger已经加载完成
func (i *HostServiceImpl) Config() {
	// Host service的子loggger
	// 封装的zap让其满足logger接口
	// 为什么要封装：
	// 		1. logger全局实例
	// 		2. logger level的动态调整，logrus不支持level动态调整
	// 		3. 加入日志轮转功能的集合
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()
}

// 返回服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

// _ import app 自动执行注册逻辑
func init() {
	// 对象注册到IOC层
	apps.RegistryImpl(impl)
}

// _ import app 自动执行注册逻辑

// 之前都是在start时候，手动把服务实现注册到ios层
// 注册hostservice的实例到IOC
// apps.HostService= impl.NewHostServiceImpl()

// mysql的驱动加载的实现方式
// sql 这个库，是一个框架，驱动是引入依赖的时候加载
// 我们把app模块，比作一个驱动，ioc比作框架
// _ import app,该app就注册到ioc层
