package AdapterDesignPattern_51

/*
对于这个模式，有⼀个经常被拿来解释它的例⼦，就是 USB 转接头充当适配器，
把两种不兼容的接⼝，通过转接变得可以⼀起工作。
原理很简单，我们再来看下它的代码实现。适配器模式有两种实现⽅式：类适配器和对象适
配器。其中，类适配器使⽤继承关系来实现，对象适配器使⽤组合关系来实现。

针对这两种实现⽅式，在实际的开发中，到底该如何选择使⽤哪⼀种呢？判断的标准主要有
两个，⼀个是 Adaptee 接⼝的个数，另⼀个是 Adaptee 和 ITarget 的契合程度。

如果 Adaptee 接⼝并不多，那两种实现⽅式都可以。
如果 Adaptee 接⼝很多，⽽且 Adaptee 和 ITarget 接⼝定义⼤部分都相同，那我们推
荐使⽤类适配器，因为 Adaptor 复⽤⽗类 Adaptee 的接⼝，⽐起对象适配器的实现⽅
式，Adaptor 的代码量要少⼀些。
如果 Adaptee 接⼝很多，⽽且 Adaptee 和 ITarget 接⼝定义⼤部分都不相同，那我们
推荐使⽤对象适配器，因为组合结构相对于继承更加灵活。

代理模式：代理模式在不改变原始类接⼝的条件下，为原始类定义⼀个代理类，主要⽬的是控制访问，⽽⾮加强功能，这是它跟装饰器模式最⼤的不同。
桥接模式：桥接模式的⽬的是将接⼝部分和实现部分分离，从⽽让它们可以较为容易、也相 对独⽴地加以改变。
装饰器模式：装饰者模式在不改变原始类接⼝的情况下，对原始类功能进⾏增强，并且⽀持 多个装饰器的嵌套使⽤。
适配器模式：适配器模式是⼀种事后的补救策略。适配器提供跟原始类不同的接⼝，⽽代理模式、装饰器模式提供的都是跟原始类相同的接⼝。
*/

type IDataLoader interface {
	Load(string) error
}

type XMLLoader struct {
}

func (x *XMLLoader) Parse(filePath string) error {
	// do action
	return nil
}

type XMLAdapterLoader struct {
	XMLLoader XMLLoader
}

func (X *XMLAdapterLoader) Load(filePath string) error {
	_ = X.XMLLoader.Parse(filePath)
	return nil
}

type JsonLoader struct {
}

type JsonAdapterLoader struct {
	JsonLoader JsonLoader
}

func (j *JsonAdapterLoader) Load(filePath string) error {
	_ = j.JsonLoader.DataLoader(filePath)
	return nil
}

func (j *JsonLoader) DataLoader(filePath string) error {
	// do action
	return nil
}
