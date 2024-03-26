package http

import (
	"github.com/Astrotrain-ali/cmdb/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/restful/response"
)

// 用于暴露Host service的接口

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	// 用户传递过来的参数进行解析，实现了一个json的unmarshal
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	// 进行接口调用
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	// 成功，把对象实例返回给API调用方
	response.Success(c.Writer, ins)
}

func (h *Handler) queryHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewQueryHostFromHTTP(c.Request)
	// 进行接口调用
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}

func (h *Handler) describeHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewDescribeHostRequestWithId(c.Param("id"))
	// 进行接口调用
	set, err := h.svc.DescribeHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}
