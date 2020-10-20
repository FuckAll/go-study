package abstract

import "fmt"

/*
	产品设计：设计一个产品，将MarkDown文件转换成HTML或者Word。
	FastDoc Soft的产品便宜，并且转换速度快，而GoodDoc Soft的产品贵，但转换效果好。
    将两个产品都集成到产品中，根据用户的选择使用。
*/

type AbstractFactory interface {
	// 创建Html文档:
	createHtml(md string) HtmlDocument
	// 创建Word文档:
	createWord(md string) WordDocument
}

// Html文档接口:
type HtmlDocument interface {
	toHtml() string
	save(path string)
}

// Word文档接口:
type WordDocument interface {
	save(path string)
}

// FastDoc HTML产品
type FastHtmlDocument struct {
}

func (f *FastHtmlDocument) toHtml() string {
	fmt.Println("FastHtmlDocument -> 生成HTML")
	return ""
}

func (f *FastHtmlDocument) save(path string) {
	fmt.Println("FastHtmlDocument -> 保存HTML")
}

// FastDoc Word产品
type FastWordDocument struct {
}

func (f *FastWordDocument) save(path string) {
	fmt.Println("FastWordDocument -> 保存Word")
}

// FastFactory
type FastFactory struct {
}

func (f *FastFactory) createHtml(md string) HtmlDocument {
	return new(FastHtmlDocument)
}

func (f *FastFactory) createWord(md string) WordDocument {
	return new(FastWordDocument)
}

// GoodDoc HTML产品
type GoodHtmlDocument struct {
}

func (g *GoodHtmlDocument) toHtml() string {
	fmt.Println("GoodHtmlDocument -> 生成HTML")
	return "生成HTML"
}

func (g *GoodHtmlDocument) save(path string) {
	fmt.Println("GoodHtmlDocument -> 保存HTML")
}

// GoodDoc Word产品
type GoodWordDocument struct {
}

func (g *GoodWordDocument) save(path string) {
	fmt.Println("GoodWordDocument -> 保存Word")
}

// GoodFactory
type GoodFactory struct {
}

func (g *GoodFactory) createHtml(md string) HtmlDocument {
	return new(GoodHtmlDocument)
}

func (g *GoodFactory) createWord(md string) WordDocument {
	return new(GoodWordDocument)
}

func createFactory(name string) AbstractFactory {
	if name == "good" {
		return new(GoodFactory)
	}
	if name == "fast" {
		return new(FastFactory)
	}
	return nil
}
