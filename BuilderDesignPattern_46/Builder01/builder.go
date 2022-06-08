package Builder01

import "errors"

/*
实际上，如果我们并不是很关⼼对象是否有短暂的⽆效状态，也不是太在意对象是否是可变的。
⽐如，对象只是⽤来映射数据库读出来的数据，那我们直接暴露 set()⽅法来设置类的成员变量值是完全没问题的。
⽽且，使⽤建造者模式来构建对象，代码实际上是有点重复 的，ResourcePoolConﬁg 类中的成员变量，
要在Builder类中重新再定义⼀遍
*/
const (
	defaultMinIdle = 1
	defaultMaxIdle = 1
	defaultMaxTotal = 1
)
type ResourcePoolConfig struct {
	name     string
	minIdle  int
	maxIdle  int
	maxTotal int
}

func newResourcePoolConfig(name string,minIdle int, maxIdle int, maxTotal int) *ResourcePoolConfig {
	return &ResourcePoolConfig{name:name,minIdle: minIdle, maxIdle: maxIdle, maxTotal: maxTotal}
}

// 用于set ResourcePoolConfig相对应的需要校验的参数，
// 将需要校验的参数先在builder进行check，然后直接
// 赋值给对应的需要返回的结构体
type ResourcePoolConfigBuilder struct {
	name     string
	minIdle  int
	maxIdle  int
	maxTotal int
}

func (rb *ResourcePoolConfigBuilder) setName(name string) {
	rb.name = name
}

func (rb *ResourcePoolConfigBuilder) setMinIdle(minIdle int) {
	rb.minIdle = minIdle
}

func (rb *ResourcePoolConfigBuilder) setMaxIdle(maxIdle int) {
	rb.maxIdle = maxIdle
}

func (rb *ResourcePoolConfigBuilder) setMaxTotal(maxTotal int) {
	rb.maxTotal = maxTotal
}
func (rb *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig,error){
	if rb.name == ""{
		return nil,errors.New("name is needed")
	}
	if rb.minIdle == 0{
		rb.minIdle = defaultMinIdle
	}
	if rb.maxIdle == 0{
		rb.maxIdle = defaultMaxIdle
	}
	if rb.maxTotal == 0{
		rb.maxTotal = defaultMaxTotal
	}
	return newResourcePoolConfig(rb.name,rb.minIdle,rb.maxIdle,rb.maxTotal), nil
}
