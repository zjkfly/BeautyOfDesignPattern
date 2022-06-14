package DecoratorDesignPattern_50

import "fmt"

/*

就拿⽐较相似的代理模式和装饰器模式来说吧，代理模式中，代理类附加
的是跟原始类⽆关的功能，⽽在装饰器模式中，**装饰器类附加的是跟原始类相关的增强功能**.

Decorator关注为对象动态的添加功能, Proxy关注对象的信息隐藏及访问控制.
Decorator体现多态性, Proxy体现封装性.

*/

type ISinger interface {
	sing() error
}

type Singer struct {
}

func (s *Singer) sing() error {
	fmt.Println("i am singing")
	return nil
}

type StarSinger struct {
	Singer ISinger
}

func (s *StarSinger) sing() error {
	// do action,such as checking micro and note singer 
	s.Singer.sing()
	return nil
}

func NewStarSinger(singer ISinger) *StarSinger {
	return &StarSinger{Singer: singer}
}




