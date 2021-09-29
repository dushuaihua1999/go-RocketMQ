package scheduler

import "go-RocketMQ/internal/namesrv/registry"

/*
	初步想设计一个调度器的manager，用于对不同策略，或者不同调度对象（如producer,consumer采取不同的调度对象）
	让调度器的创建的数量与调度策略类型的 权利交给生产者，消费者共同商榷
*/
// 先在这里设计一个调度器管理者对象的接口，具体实现可以自己决定,实现后可以用这个对调度器进行管理
type ScheManager interface {
	electOne(kind string) Scher  // 具体的选举策略可以根据情况自己实现,里面可以写具体的策略
	listAll() map[string][]Scher // string对应具体是哪个类型，如producer,consumer,broker等等
	insert(scher Scher) bool
}

type SchedulerManager struct {
	pS []Scher // producer 调度者切片
	cS []Scher // consumer 调度者切片
	bS []Scher // broker   调度器切片
}

// NewSchedulerManager pCount,cCount,bCount 分别代表对应初始化创建时用户给定的每个调度切片的初始化容量.小数据量情况下，可以传小一点
func NewSchedulerManager(pCount int, cCount int, bCount int) ScheManager {
	return &SchedulerManager{
		pS: make([]Scher, pCount),
		cS: make([]Scher, cCount),
		bS: make([]Scher, bCount),
	}
}

func (sm *SchedulerManager) electOne(kind string) Scher {
	return nil
}

func (sm *SchedulerManager) listAll() map[string][]Scher {
	return nil
}

func (sm *SchedulerManager) insert(scher Scher) bool {
	return true
}

// 对外的调度器接口
type Scher interface {
	Load() error
	Reload() error
}

// 调度器对象
type scheduler struct {
	r    *registry.Registry
	kind string
}

func (s *scheduler) Load() error {
	return nil
}

func (s *scheduler) Reload() error {
	return nil
}

func NewScheduler(r *registry.Registry) *scheduler {
	return &scheduler{
		r: r,
	}
}
