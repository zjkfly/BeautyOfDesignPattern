package BridgeDesignPattern_49

/*
Decouple an abstraction from its implementation so
that the two can vary independently。”
翻译成中⽂就是：“将抽象和实现解耦，让它们可以独⽴变化。”关于桥接模式，很多书籍、资料

===============good case=======================
举个很简单的例⼦，现在有两个纬度
Car ⻋ （奔驰、宝⻢、奥迪等）
Transmission 档位类型 （⾃动挡、⼿动挡、⼿⾃⼀体等）
按照继承的设计模式，Car是⼀个Abstract基类，假设有M个⻋品牌，N个档位⼀共要写M*N个类去描述
所有⻋和档位的结合。
⽽当我们使⽤桥接模式的话，我⾸先new⼀个具体的Car（如奔驰），再new⼀个具体的Transmission
（⽐如⾃动档）。然后奔驰.set(⼿动档)就可以了。
那么这种模式只有M+N个类就可以描述所有类型，这就是M*N的继承类爆炸简化成了M+N组合。
===============good case=======================

使用组合，而非继承,下文中的发送（sender）+消息提醒（notification）组合，防止出现继承几何级增长，
并且实现和抽象分离
*/

type ISender interface {
	send() error
}

type EmailSender struct {
	Email string
}

func (e *EmailSender) send() error {
	panic("implement me")
}

type TelephoneSender struct {
	Telephone string
}

func (t *TelephoneSender) send() error {
	panic("implement me")
}

type WechatSender struct {
	Wechat string
}

func (w *WechatSender) send() error {
	panic("implement me")
}

type INotificator interface {
	notify() error
}

type ServerNotification struct {
	TelephoneSender ISender
}

type UrgencyNotification struct {
	WechatSender ISender
}

type NormalNotification struct {
	EmailSender ISender
}
