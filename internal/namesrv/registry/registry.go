package registry

import (
	"go-RocketMQ/internal/producer"
	"go-RocketMQ/internal/consumer"
	"go-RocketMQ/internal/broker"
	"sync"
)

type Registry struct {
	prods   map[string]*producer.Producer
	cons    map[string]*consumer.Consumer
	brokers map[string]*broker.Broker
	rLock   sync.RWMutex
	scheduler *scheduler
}







