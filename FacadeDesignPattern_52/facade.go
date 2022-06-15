package FacadeDesignPattern_52

/*
门面模式解决以下问题:
1. 解决易⽤性问题
2. 解决性能问题
3. 解决分布式事务问题
*/

type ILogin interface {
	Login() error
}
type IRegister interface {
	Register() error
}

type Facade interface {
	RegisterOrLogin() error
}

type User struct {
}

func (u *User) RegisterOrLogin() error {
	err := u.Login();
	if err != nil {
		return err
	}else {
		return u.Register()
	}
}

func (u *User) Register() error {
	panic("implement me")
}

func (u *User) Login() error {
	panic("implement me")
}
