package visitor

import "fmt"

/*
	访问者模式：允许一个或者多个操作应用到一组对象上，解耦操作和对象本身。
	每个对象accept方法接收访问者的处理。
	在增加访问者的时候，只需要增加新的访问者类即可。符合单一职责，开闭原则，并且解耦对象和对象的操作，降低对象的复杂度
*/

type ResourceFile interface {
	accept(visitor Visitor)
}

type PDFFile struct {
	filePath string
}

func (pdf *PDFFile) accept(visitor Visitor) {
	visitor.visit(pdf)
}

type WordFile struct {
}

func (word *WordFile) accept(visitor Visitor) {
	visitor.visit(word)
}

type PPTFile struct {
}

func (ppt *PPTFile) accept(visitor Visitor) {
	visitor.visit(ppt)
}

type Visitor interface {
	visit(file ResourceFile)
}

type Extractor struct {
}

func (e *Extractor) visit(file ResourceFile) {
	fmt.Println("Extractor complete.")
}

type Compressor struct {
}

func (c *Compressor) visit(file ResourceFile) {
	fmt.Println("Compressor complete.")
}
