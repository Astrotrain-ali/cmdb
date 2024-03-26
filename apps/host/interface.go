package host

import (
	"context"
)

// host app service 的接口定义
type Service interface {
	// 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// 查询主机详情
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	// 主机更新
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// 主机删除, 比如前端需要打印当前删除主机的Ip或者其他信息
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}
