package registry

import (
	"go-RocketMQ/conf"
	"go-RocketMQ/internal/broker"
	"go-RocketMQ/internal/consumer"
	"go-RocketMQ/internal/namesrv/scheduler"
	"go-RocketMQ/internal/producer"
	"sync"
)

/*
	接口接受对象时要求必须时指针对象，也就是可以理解为接口本身就是指针
*/
type Registry struct {
	prods     map[string]*producer.Producer // 生产者注册集合
	cons      map[string]*consumer.Consumer // 消费者注册集合
	brokers   map[string]*broker.Broker     // Broker注册集合
	rLock     sync.RWMutex                  // 用于控制对上面三个的map执行并发访问的控制
	scheduler scheduler.Scher               // 调度选用可用Broker或者对应topic的Broker返回给消费者或者生产者
}

func NewRegistry(conf *conf.Config) (r *Registry) {
	r = &Registry{
		prods:   make(map[string]*producer.Producer),
		cons:    make(map[string]*consumer.Consumer),
		brokers: make(map[string]*broker.Broker),
	}
	// 将registry注册表传给调度器
	r.scheduler = scheduler.NewScheduler(r)
	// 初次加载调度区
	r.scheduler.Load()
	go r.scheduler.Reload()
	//go r.proc()
	return r
}
