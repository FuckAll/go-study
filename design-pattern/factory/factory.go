package factory

import "sync"

/*
	工厂模式:对于复杂的创建过程，每个类自身有简单工厂模式。相当于简单工厂的工厂方法。
*/

type IRuleConfigParserFactory interface {
	createParser() IRuleConfigParse
}

type JSONRuleConfigParserFactory struct {
}

func (j *JSONRuleConfigParserFactory) createParser() IRuleConfigParse {
	// 假设这里创建过程复杂，是个简单工厂模式
	return new(JSONRuleConfigParse)
}

type XMLRuleConfigParserFactory struct {
}

func (j *XMLRuleConfigParserFactory) createParser() IRuleConfigParse {
	return new(XMLRuleConfigParse)
}

type YAMLRuleConfigParserFactory struct {
}

func (j *YAMLRuleConfigParserFactory) createParser() IRuleConfigParse {
	return new(XMLRuleConfigParse)
}

type PropertiesRuleConfigParserFactory struct {
}

func (j *PropertiesRuleConfigParserFactory) createParser() IRuleConfigParse {
	return new(XMLRuleConfigParse)
}

var cachedFactoryParsers = map[string]IRuleConfigParserFactory{}
var onceFactory sync.Once

// RuleConfigParserFactory 工厂模式
func RuleConfigParserFactory(fileType string) IRuleConfigParserFactory {
	onceFactory.Do(func() {
		cachedFactoryParsers["json"] = new(JSONRuleConfigParserFactory)
		cachedFactoryParsers["xml"] = new(XMLRuleConfigParserFactory)
		cachedFactoryParsers["yaml"] = new(YAMLRuleConfigParserFactory)
		cachedFactoryParsers["properties"] = new(PropertiesRuleConfigParserFactory)
	})
	return cachedFactoryParsers[fileType]
}
