package http

import (
	"github.com/Astrotrain-ali/cmdb/apps"
	"github.com/gin-gonic/gin"
)

// 面向接口，真正Service的实现，在服务实例化的时候传递进行，也就是（CLI）Start时候
var handler = &Handler{}

// 通过写一个实例类，把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类，会实现gin的http handler
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {

	// 从IOC里面获取HostService的实例对象
	h.svc = apps.GetImpl(host.AppName).(host.Service)
}

// 完成了 Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
	r.GET("/hosts", h.queryHost)
	r.GET("/hosts/:id", h.describeHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

// 完成Http Handler注册
func init() {
	apps.RegistryGin(handler)
}
