package FactoryDesignPattern_44

import "errors"

/*

抽象工厂

Q: 过多的类 也会让系统难维护。这个问题该怎么解决呢？
A: 抽象⼯⼚就是针对这种⾮常特殊的场景⽽诞⽣的。
   我们可以让⼀个⼯⼚负责创建多个不同类 型的对象（IRuleConﬁgParser、ISystemConﬁgParser 等）
*/

type IRuleConfigParser interface {
	parse() RuleConfig
}
type ISystemConfigParser interface {
	parse() RuleConfig
}
type JsonRuleConfigParser struct {
}
type JsonSystemConfigParser struct {
}

type IConfigParserFactory interface {
	createRuleParser() IRuleConfigParser
	createSystemParser() ISystemConfigParser
}

type JsonConfigParserFactory struct {
}

func (j *JsonConfigParserFactory) createRuleParser() IRuleConfigParser {
	return NewJsonRuleConfigParser()
}
func (j *JsonConfigParserFactory) createSystemParser() ISystemConfigParser {
	return NewJsonSystemConfigParser()
}

func NewJsonRuleConfigParser() IRuleConfigParser {
	return &JsonRuleConfigParser{}
}

func NewJsonSystemConfigParser() ISystemConfigParser {
	return &JsonSystemConfigParser{}
}

func (j *JsonRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

func (j *JsonSystemConfigParser) parse() RuleConfig {
	panic("implement me")
}

type XmlRuleConfigParser struct {
}

func NewXmlRuleConfigParser() IRuleConfigParser {
	return &XmlRuleConfigParser{}
}

func (x *XmlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type YamlRuleConfigParser struct {
}

func NewYamlRuleConfigParser() IRuleConfigParser {
	return &YamlRuleConfigParser{}
}

func (y *YamlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type PropertiesRuleConfigParser struct {
}

func NewPropertiesRuleConfigParser() IRuleConfigParser {
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
	cachedParsers map[string]IRuleConfigParser
}

func NewRuleConfigParserFactory() *RuleConfigParserFactory {
	return &RuleConfigParserFactory{
		cachedParsers: map[string]IRuleConfigParser{
			"json":       NewJsonRuleConfigParser(),
			"xml":        NewXmlRuleConfigParser(),
			"yaml":       NewYamlRuleConfigParser(),
			"properties": NewPropertiesRuleConfigParser(),
		},
	}
}
func (r *RuleConfigParserFactory) createParser(configFmt string) IRuleConfigParser {
	return r.cachedParsers[configFmt]
}
