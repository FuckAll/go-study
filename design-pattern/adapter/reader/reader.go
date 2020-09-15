package reader

import "fmt"

type Reader struct {
	book              IBook
	currentPageNumber int
}

func (r *Reader) InitBook(book IBook) {
	r.book = book
	r.currentPageNumber = 0
}

func (r *Reader) OpenFile(filePath string) {
	r.book.ParseFile(filePath)
}

func (r *Reader) CloseFile(file string) {
	fmt.Printf("close file %s\n", file)
}

func (r *Reader) PrePage() Page {
	if r.currentPageNumber <= 0 {
		return Page{}
	}
	r.currentPageNumber--
	return r.book.GetPage(r.currentPageNumber)
}

func (r *Reader) NextPage() Page {
	if r.book.GetPageCount() >= r.currentPageNumber {
		return Page{}
	}
	r.currentPageNumber++
	return r.book.GetPage(r.currentPageNumber)
}

func (r *Reader) GotoPage(pageNumber int) Page {
	if pageNumber < 0 && pageNumber > r.book.GetPageCount() {
		return Page{}
	}

	r.currentPageNumber = pageNumber
	return r.book.GetPage(pageNumber)
}
