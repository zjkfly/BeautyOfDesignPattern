package FactoryDesignPattern_44

import "errors"

/*
⼀般情况下，⼯⼚模式分为三种更加细分的类型：简单⼯⼚、⼯⼚⽅法和抽象⼯⼚。
不过， 在GoF 的《设计模式》⼀书中，它将简单⼯⼚模式看作是⼯⼚⽅法模式的⼀种特例，
所以 ⼯⼚模式只被分成了⼯⼚⽅法和抽象⼯⼚两类。实际上，前⾯⼀种分类⽅法更加常⻅，
所 以，在今天的讲解中，我们沿⽤第⼀种分类⽅法。

简单⼯⼚模式的代码实现中，有多处 if分⽀判断逻辑，违背开闭原则，但权衡扩展性和可读性。
这样的代码实现在⼤多数情况下（⽐如，不需要频繁地添加 parser，也没有太多的parser）是没有问题的。


⼯⼚模式需要额外创建诸多 Factory 类，也会增加代码的复杂性，⽽且，每个 Factory类只是做简单的new 操作，
功能⾮常单 薄（只有⼀⾏代码），也没必要设计成独⽴的类，所以，在这个应⽤场景下，简单⼯⼚模式简单好⽤，
⽐⼯⼚⽅法模式更加合适

*/

type IParse interface {
	parse() RuleConfig
}
type JsonRuleConfigParser struct {
}

func NewJsonRuleConfigParser() IParse {
	return &JsonRuleConfigParser{}
}

func (j *JsonRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type XmlRuleConfigParser struct {
}

func NewXmlRuleConfigParser() IParse {
	return &XmlRuleConfigParser{}
}

func (x *XmlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type YamlRuleConfigParser struct {
}

func NewYamlRuleConfigParser() IParse {
	return &YamlRuleConfigParser{}
}

func (y *YamlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type PropertiesRuleConfigParser struct {
}

func NewPropertiesRuleConfigParser() IParse {
	return &PropertiesRuleConfigParser{}
}

func (p *PropertiesRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type RuleConfigSource struct {
}

type RuleConfig struct {
}

func (r *RuleConfigSource) Load(ruleConfigFilePath string) (*RuleConfig, error) {
	ruleConfigFileExtension := getFileExtension(ruleConfigFilePath)
	ruleConfigParserFactory := NewRuleConfigParserFactory()
	parser := ruleConfigParserFactory.createParser(ruleConfigFileExtension)
	if parser == nil {
		return nil, errors.New("not found parser")
	}
	ruleConfig := parser.parse()
	return &ruleConfig, nil
}

func getFileExtension(filePath string) string {
	//...解析⽂件名获取扩展名，⽐如rule.json，返回json
	return "json"
}

type RuleConfigParserFactory struct {
	cachedParsers map[string]IParse
}

func NewRuleConfigParserFactory() *RuleConfigParserFactory {
	return &RuleConfigParserFactory{
		cachedParsers: map[string]IParse{
			"json":       NewJsonRuleConfigParser(),
			"xml":        NewXmlRuleConfigParser(),
			"yaml":       NewYamlRuleConfigParser(),
			"properties": NewPropertiesRuleConfigParser(),
		},
	}
}
func (r *RuleConfigParserFactory) createParser(configFmt string) IParse {
	return r.cachedParsers[configFmt]
}
