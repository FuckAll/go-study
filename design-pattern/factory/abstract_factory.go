package factory

import "sync"

/*
	抽象工厂：简单工厂和工厂模式都是创建一种类型的类，而抽象工厂模式可以创建多种类型的类
*/

type IAbstractFactory interface {
	createRuleConfigParse() IRuleConfigParse
	createSystemRuleConfigParse() ISystemRuleConfigParser
}

type ISystemRuleConfigParser interface {
	systemParse()
}

type JSONSystemRuleConfigParse struct {
}

func (j *JSONSystemRuleConfigParse) systemParse() {}

type XMLSystemRuleConfigParse struct {
}

func (j *XMLSystemRuleConfigParse) systemParse() {}

type YAMLSystemRuleConfigParse struct {
}

func (j *YAMLSystemRuleConfigParse) systemParse() {}

type PropertiesSystemRuleConfigParse struct {
}

func (j *PropertiesSystemRuleConfigParse) systemParse() {}

type JSONParserFactory struct {
}

func (j *JSONParserFactory) createRuleConfigParse() IRuleConfigParse {
	// 假设这里创建过程复杂，是个简单工厂模式
	return new(JSONRuleConfigParse)
}

func (j *JSONParserFactory) createSystemRuleConfigParse() ISystemRuleConfigParser {
	return new(JSONSystemRuleConfigParse)
}

type XMLParserFactory struct {
}

func (j *XMLParserFactory) createRuleConfigParse() IRuleConfigParse {
	// 假设这里创建过程复杂，是个简单工厂模式
	return new(XMLRuleConfigParse)
}

func (j *XMLParserFactory) createSystemRuleConfigParse() ISystemRuleConfigParser {
	return new(XMLSystemRuleConfigParse)
}

type YAMLParserFactory struct {
}

func (j *YAMLParserFactory) createRuleConfigParse() IRuleConfigParse {
	// 假设这里创建过程复杂，是个简单工厂模式
	return new(YAMLRuleConfigParse)
}

func (j *YAMLParserFactory) createSystemRuleConfigParse() ISystemRuleConfigParser {
	return new(YAMLSystemRuleConfigParse)
}

type PropertiesParserFactory struct {
}

func (j *PropertiesParserFactory) createRuleConfigParse() IRuleConfigParse {
	// 假设这里创建过程复杂，是个简单工厂模式
	return new(PropertiesRuleConfigParse)
}

func (j *PropertiesParserFactory) createSystemRuleConfigParse() ISystemRuleConfigParser {
	return new(PropertiesSystemRuleConfigParse)
}

var cachedAbstractFactoryParsers = map[string]IAbstractFactory{}
var onceAbstract sync.Once

// RuleConfigParseAbstractFactory 抽象工厂模式
func RuleConfigParseAbstractFactory(fileType string) IAbstractFactory {
	onceAbstract.Do(func() {
		cachedAbstractFactoryParsers["json"] = new(JSONParserFactory)
		cachedAbstractFactoryParsers["xml"] = new(XMLParserFactory)
		cachedAbstractFactoryParsers["yaml"] = new(YAMLParserFactory)
		cachedAbstractFactoryParsers["properties"] = new(PropertiesParserFactory)
	})
	return cachedAbstractFactoryParsers[fileType]
}
