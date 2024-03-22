package conf_test

import (
	"os"
	"testing"

	"github.com/Astrotrain-ali/cmdb/conf"
	"github.com/stretchr/testify/assert"
)

// 测试从配置文件加载配置
func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("demo", conf.C().App.Name)
	}
}

// 测试从环境变量加载配置
func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MYSQL_DATABASE", "unit_test")
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
	}
}

// 测试从配置文件加载配置
func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}

}
