package FactoryDesignPattern_44

import "errors"

type ParseInterface interface {
	parse() RuleConfig
}
type JsonRuleConfigParser struct {
}

func NewJsonRuleConfigParser() ParseInterface {
	return &JsonRuleConfigParser{}
}

func (j *JsonRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type XmlRuleConfigParser struct {
}

func NewXmlRuleConfigParser() ParseInterface {
	return &XmlRuleConfigParser{}
}

func (x *XmlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type YamlRuleConfigParser struct {
}

func NewYamlRuleConfigParser() ParseInterface {
	return &YamlRuleConfigParser{}
}

func (y *YamlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type PropertiesRuleConfigParser struct {
}

func NewPropertiesRuleConfigParser() ParseInterface {
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
	cachedParsers map[string]ParseInterface
}

func NewRuleConfigParserFactory() *RuleConfigParserFactory {
	return &RuleConfigParserFactory{
		cachedParsers: map[string]ParseInterface{
			"json":       NewJsonRuleConfigParser(),
			"xml":        NewXmlRuleConfigParser(),
			"yaml":       NewYamlRuleConfigParser(),
			"properties": NewPropertiesRuleConfigParser(),
		},
	}
}
func (r *RuleConfigParserFactory) createParser(configFmt string) ParseInterface {
	return r.cachedParsers[configFmt]
}
