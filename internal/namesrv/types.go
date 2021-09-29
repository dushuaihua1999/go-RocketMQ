package namesrv

import (
	"go-RocketMQ/conf"
	"go-RocketMQ/internal/namesrv/registry"
	"go-RocketMQ/internal/types"
	"log"
	"sync"
)

//存放namesrv的数据结构

type NameSrv struct {
	c      *conf.Config
	status types.Status  // 公共状态属性
	sLock  sync.RWMutex	 // 读写锁
	nlog   log.Logger	 // 自带的日志对象,可以自定义实现,我用的是b站kartos的log
	register *registry.Registry
}


// NewNameSrv 新建一个NameSrv
func NewNameSrv() *NameSrv {

}


