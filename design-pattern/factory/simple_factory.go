package factory

import "sync"

/*
	简单工厂模式
*/
type IRuleConfigParse interface {
	parse()
}

type JSONRuleConfigParse struct {
}

func (j *JSONRuleConfigParse) parse() {}

type XMLRuleConfigParse struct {
}

func (j *XMLRuleConfigParse) parse() {}

type YAMLRuleConfigParse struct {
}

func (j *YAMLRuleConfigParse) parse() {}

type PropertiesRuleConfigParse struct {
}

func (j *PropertiesRuleConfigParse) parse() {}

var cachedSimpleFactoryParsers = map[string]IRuleConfigParse{}
var once sync.Once

// RuleConfigParseSimpleFactory 简单工厂模式
func RuleConfigParseSimpleFactory(fileType string) IRuleConfigParse {
	once.Do(func() {
		cachedSimpleFactoryParsers["json"] = new(JSONRuleConfigParse)
		cachedSimpleFactoryParsers["xml"] = new(XMLRuleConfigParse)
		cachedSimpleFactoryParsers["yaml"] = new(YAMLRuleConfigParse)
		cachedSimpleFactoryParsers["properties"] = new(PropertiesRuleConfigParse)
	})
	return cachedSimpleFactoryParsers[fileType]
}
