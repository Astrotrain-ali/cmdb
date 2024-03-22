package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Astrotrain-ali/cmdb/apps/host"
	"github.com/Astrotrain-ali/cmdb/apps/host/impl"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
)

var (
	service host.Service
)

func TestCreste(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "hangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init() {
	// 需要初始化全局logger
	// 为什么不设计为默认打印，因为性能
	zap.DevelopmentSetup()
	// host service 的具体实现
	service = impl.NewHostServiceImpl()
}
