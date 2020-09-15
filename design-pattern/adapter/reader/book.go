package reader

type IBook interface {
	ParseFile(path string) bool
	GetCatalogue() Catalogue
	GetPageCount() int
	GetPage(pageNumber int) Page
}

type Page struct {
}

type Catalogue struct {
}

type TxtBook struct {
}

func (t *TxtBook) ParseFile(path string) bool {
	return true
}

func (t *TxtBook) GetCatalogue() Catalogue {
	return Catalogue{}
}

func (t *TxtBook) GetPageCount() int {
	return 100
}

func (t *TxtBook) GetPage(pageNumber int) Page {
	return Page{}
}

type EpubBook struct {
}

func (e *EpubBook) ParseFile(path string) bool {
	return true
}

func (e *EpubBook) GetCatalogue() Catalogue {
	return Catalogue{}
}

func (e *EpubBook) GetPageCount() int {
	return 100
}

func (e *EpubBook) GetPage(pageNumber int) Page {
	return Page{}
}

type ThirdPdf struct {
}

type Outline struct {
}

type PdfPage struct {
}

func (t *ThirdPdf) Open(filePath string) {

}

func (t *ThirdPdf) GetOutline() Outline {
	return Outline{}
}

func (t *ThirdPdf) PageSize() int {
	return 100
}

func (t *ThirdPdf) Page(pageNumber int) PdfPage {
	return PdfPage{}
}

// 适配器模式，将第三方的pdf适配成Reader能够使用的IBook
type PdfAdapterBook struct {
	thirdPdf ThirdPdf
}

func (p *PdfAdapterBook) ParseFile(path string) bool {
	p.thirdPdf.Open(path)
	return true
}

func (p *PdfAdapterBook) GetCatalogue() Catalogue {
	p.thirdPdf.GetOutline()
	// outline 转换 catalogue
	return Catalogue{}
}

func (p *PdfAdapterBook) GetPageCount() int {
	return p.thirdPdf.PageSize()
}

func (p *PdfAdapterBook) GetPage(pageNumber int) Page {
	p.thirdPdf.Page(pageNumber)
	// PdfPage 转换 page
	return Page{}
}
