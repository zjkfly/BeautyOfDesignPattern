package FactoryDesignPattern_44

/*
工厂模式实践
*/

type IParse interface {
	parse() RuleConfig
}
type IRuleConfigParserFactory interface {
	createParser() IParse
}
type JsonRuleConfigParserFactory struct {
}

func NewJsonRuleConfigParserFactory() IRuleConfigParserFactory {
	return &JsonRuleConfigParserFactory{}
}

func (j *JsonRuleConfigParserFactory) createParser() IParse {
	return NewJsonRuleConfigParser()
}

type JsonRuleConfigParser struct {
}

func NewJsonRuleConfigParser() *JsonRuleConfigParser {
	return &JsonRuleConfigParser{}
}

func (j *JsonRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}

type XmlRuleConfigParserFactory struct {
}

func NewXmlRuleConfigParserFactory() IRuleConfigParserFactory {
	return &XmlRuleConfigParserFactory{}
}

func (x *XmlRuleConfigParserFactory) createParser() IParse {
	return NewXmlRuleConfigParser()
}

type XmlRuleConfigParser struct {
}

func NewXmlRuleConfigParser() IParse {
	return &XmlRuleConfigParser{}
}

func (x *XmlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}
type YamlRuleConfigParserFactory struct {
}

func NewYamlRuleConfigParserFactory() IRuleConfigParserFactory {
	return &YamlRuleConfigParserFactory{}
}

func (y *YamlRuleConfigParserFactory) createParser() IParse {
	return NewYamlRuleConfigParser()
}
type YamlRuleConfigParser struct {
}

func NewYamlRuleConfigParser() IParse {
	return &YamlRuleConfigParser{}
}

func (y *YamlRuleConfigParser) parse() RuleConfig {
	panic("implement me")
}
type PropertiesRuleConfigParserFactory struct {
}

func NewPropertiesRuleConfigParserFactory() IRuleConfigParserFactory {
	return &PropertiesRuleConfigParserFactory{}
}

func (p *PropertiesRuleConfigParserFactory) createParser() IParse {
	return NewPropertiesRuleConfigParser()
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
	ruleConfigParserFactories := NewRuleConfigParserFactoryMap()
	factory := ruleConfigParserFactories.getParserFactory(ruleConfigFileExtension)
	ruleConfig := factory.createParser().parse()
	return &ruleConfig, nil
}

func getFileExtension(filePath string) string {
	//...解析⽂件名获取扩展名，⽐如rule.json，返回json
	return "json"
}

type RuleConfigParserFactory struct {
	cachedParsers map[string]IParse
}

type RuleConfigParserFactories struct {
	cachedFactories map[string]IRuleConfigParserFactory
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



func (r *RuleConfigParserFactories)getParserFactory(configFmt string) IRuleConfigParserFactory {
	return r.cachedFactories[configFmt]
}
func NewRuleConfigParserFactoryMap() *RuleConfigParserFactories {
	return &RuleConfigParserFactories{
		cachedFactories: map[string]IRuleConfigParserFactory{
			"json":       NewJsonRuleConfigParserFactory(),
			"xml":        NewXmlRuleConfigParserFactory(),
			"yaml":       NewYamlRuleConfigParserFactory(),
			"properties": NewPropertiesRuleConfigParserFactory(),
		},
	}
}
