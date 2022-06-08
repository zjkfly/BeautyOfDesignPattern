package ProxyDesignPattern_48

type IUser interface {
	login()
	register()
}

type UserController struct {
}

func (u *UserController) login() {
	panic("implement me")
}

func (u *UserController) register() {
	panic("implement me")
}

// di
type ProxyUserController struct {
	UserController UserController
}

func (p *ProxyUserController) login() {
	p.UserController.login()
	// do something 非业务性操作
}

func (p *ProxyUserController) register() {
	p.UserController.register()
	// do something 非业务性操作
}
