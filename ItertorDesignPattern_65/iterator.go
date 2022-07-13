package main

/*
这⾥说的“集合对象”也可以叫“容器”“聚合对象”，实际上就是包含⼀组对象的对象，
⽐如数组、链表、树、图、跳表。
迭代器模式将集 合对象的遍历操作从集合类中拆分出来，放到迭代器类中，
让两者的职责更加单⼀。
*/

type Iterator interface {
	HasNext() bool
	GetNext() error
	CurrentItem() interface{}
}

type User struct {
}

type UserList struct {
	Users []User
}

func NewUserList(users []User) *UserList {
	return &UserList{Users: users}
}

func (u UserList) GetIterator() Iterator {
	return &UserList{Users: nil}
}

func (receiver *UserList) GetNext() error {
	panic("implement me")
}

func (receiver *UserList) CurrentItem() interface{} {
	panic("implement me")
}

func (receiver *UserList) HasNext() bool {
	return true
}
