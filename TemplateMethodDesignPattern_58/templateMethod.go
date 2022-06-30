package main

/*
模板⽅法模式在⼀个⽅法中定义⼀个算法⻣架，并将某些步骤推迟到⼦类
中实现。模板⽅法模式可以让⼦类在不改变算法整体结构的情况下，重新定义算法中的某些步奏。

模板方法的主要两个左右：1.复用 2.扩展
*/

/*
IOtp是一次性密码的例子
1生成随机的 n 位数字。
2在缓存中保存这组数字以便进行后续验证。
3准备内容。
4发送通知。
5发布。
*/

type IOtp interface {
	genRandomN(n int) int
	saveCache(n int) error
	prepareMsg(n int) string
	sendMsg(msg string) error
	publish(msg string)
}

type Otp struct {
	iOtp IOtp
}
/*
定义一个由固定数量的方法组成的基础模板算法。
这就是我们的模板方法。然后我们将实现每一个步骤方法，但不会改变模板方法。

也可以不使用注入，直接func返回，但是没有必要
*/
func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomN(otpLength)
	o.iOtp.saveCache(otp)
	message := o.iOtp.prepareMsg(otp)
	err := o.iOtp.sendMsg(message)
	if err != nil {
		return err
	}
	o.iOtp.publish(message)
	return nil
}

type SmsOtp struct {

}
type EmailOtp struct {

}

func (e *EmailOtp) genRandomN(n int) int {
	panic("implement me")
}

func (e *EmailOtp) saveCache(n int) error {
	panic("implement me")
}

func (e *EmailOtp) prepareMsg(n int) string {
	panic("implement me")
}

func (e *EmailOtp) sendMsg(msg string) error {
	panic("implement me")
}

func (e *EmailOtp) publish(msg string) {
	panic("implement me")
}

func (s *SmsOtp) genRandomN(n int) int {
	panic("implement me")
}

func (s *SmsOtp) saveCache(n int) error {
	panic("implement me")
}

func (s *SmsOtp) prepareMsg(n int) string {
	panic("implement me")
}

func (s *SmsOtp) sendMsg(msg string) error {
	panic("implement me")
}

func (s *SmsOtp) publish(msg string) {
	panic("implement me")
}

func main() {
	sms := &SmsOtp{}
	smsOtp := Otp{
		iOtp: sms,
	}
	email := &EmailOtp{}
	emailOtp := Otp{
		iOtp: email,
	}
	smsOtp.genAndSendOTP(6)
	emailOtp.genAndSendOTP(6)
}

