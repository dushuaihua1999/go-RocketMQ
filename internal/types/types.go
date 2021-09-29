package types

type Status struct {
	Name string
	Role string
	IsRunning bool
	Others map[string]interface{}
}



// producer,consumer,namesrv,broker公共的接口与类型
// Participant 消息队列的参与者接口，让里面的相关对象都实现这个接口，不让在方法调用的时候暴露细节
//
type Participant interface {
	GetStatus() Status
}
