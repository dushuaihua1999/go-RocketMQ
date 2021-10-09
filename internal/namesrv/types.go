package namesrv

import (
	"go-RocketMQ/conf"
	"go-RocketMQ/internal/namesrv/registry"
	"go-RocketMQ/internal/registry_public"
	"go-RocketMQ/internal/types"
	"log"
	"sync"
)

/*
	1.Topic路由注册中心 -- 支持Broker的动态注册与发现

	2.连个主要功能:
		2.1  Broker管理，NameServer接受Broker集群的注册信息并且保存下来作为路由信息的基本数据。
			然后提供心跳检测机制，检查Broker是否还存活
		2.2  路由信息管理，每个NameServer将保存关于Broker集群的整个路由信息和用于客户端查询的
             队列信息。然后Producer和Conumser通过NameServer就可以知道整个Broker集群的路由信息，
			从而进行消息的投递和消费

	3.NameServer通常也是集群的方式部署，各实例间相互不进行信息通讯

	4.Broker是向每一台NameServer注册自己的路由信息，所以每一个NameServer实例上面都保存一份完整
	  的路由信息。当某个NameServer因某种原因下线了，Broker仍然可以向其它NameServer同步其路由信息，
	  Producer,Consumer仍然可以动态感知Broker的路由的信息。
*/

//存放namesrv的数据结构
type NameSrv struct {
	c        *conf.Config
	status   types.Status // 公共状态属性
	sLock    sync.RWMutex // 读写锁
	nlog     log.Logger   // 自带的日志对象,可以自定义实现,我用的是b站kartos的log
	register *registry.Registry
	instance *registry_public.ServiceInstance
}

// NewNameSrv 新建一个NameSrv
func NewNameSrv(c *conf.Config) *NameSrv {
	n := &NameSrv{
		c:        c,
		register: registry.NewRegistry(c),
	}
	return n
}

// NameSrv在监听连接时，客户端的连接方式有两种，
//1.一种是使用同一种proto文件
//2.一种式Producer使用Producer的proto文件，
//Consumer使用Consumer的proto，Broker使用Broker的proto的
//第一种就要考虑同意传递的信息，通过统一proto不同service来区分不同的丽娜姐对象双方
/*
  例如：
	SendMsgService    -->  producer与broker之间的服务
	ReceiveMsgService -->  broker与consumer之间的服务
	RegisterService   -->  broker向namesrv注册以及心跳检测的服务
	acquireRegistry   -->  consumer,producer向namresrv请求注册信息的服务
*/
//至于负载均衡，就在producer与consumer那里做
func (n *NameSrv) Start() {

}
