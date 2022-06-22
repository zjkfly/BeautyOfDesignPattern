package ObserverDesignPattern_56

/*
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
