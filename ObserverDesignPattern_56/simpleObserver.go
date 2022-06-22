package ObserverDesignPattern_56

/*
观察者模式的应⽤场景⾮常⼴泛，⼩到代码层⾯的解耦，⼤到架构层⾯的系统解耦，
再或者 ⼀些产品的设计思路，都有这种模式的影⼦，⽐如，邮件订阅、RSS Feeds，
本质上都是 观察者模式。

设计模式要⼲的事情就 是解耦。
创建型模式是将创建和使⽤代码解耦，
结构型模式是将不同功能代码解耦，
⾏为型模式是将不同的⾏为代码解耦，
具体到观察者模式，它是将观察者和被观察者代码解耦。

实现了订阅和多种类型消费者的关系
*/

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update() error
}

type Subject struct {
	Observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.Observers = append(s.Observers, observer)
}

func (s *Subject) Remove(observer IObserver) {
	flagIdx := 0
	for idx, o := range s.Observers {
		if observer == o {
			flagIdx = idx
			break
		}
	}
	s.Observers = append(s.Observers[:flagIdx], s.Observers[flagIdx:]...)
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.Observers {
		o.Update()
	}
}

type ObserverTypeAAA struct {
}

func (o *ObserverTypeAAA) Update() error {
	//TODO implement me
	panic("implement me")
}

type ObserverTypeBBB struct {
}

func (o *ObserverTypeBBB) Update() error {
	//TODO implement me
	panic("implement me")
}
