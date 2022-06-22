package ObserverDesignPattern_56

import (
	"fmt"
	"reflect"
	"sync"
)

// 单体系统可以使用eventBus来进行系统内部的生产消费，而不需要
// 使用MQ来减少网络的IO操作。

// Bus Bus
type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// AsyncEventBus 异步事件总线
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

// NewAsyncEventBus new
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

// Subscribe 订阅
func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := bus.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handlers[topic] = handler

	return nil
}

// Publish 发布
// 这里异步执行，并且不会等待返回结果
func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}
