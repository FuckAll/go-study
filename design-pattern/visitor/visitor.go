package visitor

import "fmt"

/*
	访问者模式：允许一个或者多个操作应用到一组对象上，解耦操作和对象本身。
	每个对象accept方法接收访问者的处理。
	在增加访问者的时候，只需要增加新的访问者类即可。符合单一职责，开闭原则，并且解耦对象和对象的操作，降低对象的复杂度
	缺点：违反LOD迪米特法则，违反了高内聚，对象本身的变化比较困难。
	例如：实现文件的提取、压缩功能；目前有各种的文件类型，pdf、word、ppt，通过外部的访问者提取类和压缩类，实现提取压缩的功能。
*/

// 资源类要能够接受范访问者
type ResourceFile interface {
	accept(visitor Visitor)
}

type PDFFile struct {
	filePath string
}

// 让访问者能够访问类b本身
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

// 作为访问者，能够处理需要被处理的对象本身
func (e *Extractor) visit(file ResourceFile) {
	fmt.Println("Extractor complete.")
}

type Compressor struct {
}

func (c *Compressor) visit(file ResourceFile) {
	fmt.Println("Compressor complete.")
}
