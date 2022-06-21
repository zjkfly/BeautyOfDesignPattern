package ComposeDesignPattern

/*
Compose objects into tree structure to represent part-whole
hierarchies.Composite lets client treat individual objects and
compositions of objects uniformly.

将⼀组对象组织（Compose）成树形结构，以表示⼀种“部分 - 整体”的
层次结构。组合让客户端（在很多设计模式书籍中，“客户端”代指代码的使⽤者。）可以
统⼀单个对象和组合对象的处理逻辑。

个人理解：组合模式，将⼀组对象组织成树形结构，将单个对象和组合对象都看做树中的节点，以统⼀
处理逻辑，并且它利⽤树形结构的特点，递归地处理每个⼦树，依次简化代码实现。使⽤组
合模式的前提在于，你的业务场景必须能够表示成树形结构。所以，组合模式的应⽤场景也
⽐较局限，它并不是⼀种很常⽤的设计模式。

***本质上面向接口编程***
*/

type IHumanResource interface {
	CountSalary() int64
}

type Employee struct {
	Id     int64
	Salary int64
}

func (e *Employee) CountSalary() int64 {
	return e.Salary
}

type Department struct {
	Id            int64
	Employees     []Employee
	SubDepartment []Department
}

func (d *Department) CountSalary() int64 {
	var allSalary int64
	for _,e := range d.Employees{
		allSalary+= e.CountSalary()
	}
	for _,d := range d.SubDepartment{
		allSalary+= d.CountSalary()
	}
	return allSalary
}

func (d *Department) AddEmployee(e Employee) error {
	d.Employees = append(d.Employees, e)
	return nil
}
func (d *Department) AddSubDepartment(dd Department) error {
	d.SubDepartment = append(d.SubDepartment, dd)
	return nil
}
