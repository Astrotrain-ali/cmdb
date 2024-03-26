package impl

import (
	"context"
	"fmt"

	"github.com/Astrotrain-ali/cmdb/apps/host"
)

// 完成对象和数据库之间的转换

// 把Host对象保存到数据库内,数据的一致性
func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {

	var (
		err error
	)

	// 把数据入库到resource表和host表
	// 一次需要往2个表录入数据，我们需要2个操作 要么都成功,要么都失败，事务的逻辑
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error,%s", err)
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则commit事务
	// 2. 有报错，则Rollback事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error,%s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error,%s", err)
			}
		}
	}()

	// 插入Resource 数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		i.l.Error("insert Resource error,%s", err)
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		i.l.Error("exec insert Resource error,%s", err)
	}

	// 插入Describe数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		i.l.Error("insert Describe error,%s", err)
	}
	defer dstmt.Close()

	_, err = dstmt.ExecContext(ctx,
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		i.l.Error("exec insert Describe error,%s", err)
	}

	return err
}
