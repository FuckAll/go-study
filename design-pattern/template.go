package design

import "fmt"

// Downloader ...
type Downloader interface {
	Download() error
}

type implement interface {
	download()
	save()
}

type template struct {
	implement
	uri string
}

func (t *template) Download() error {
	fmt.Println("begin download...")
	t.implement.download()
	fmt.Println("begin save...")
	t.implement.save()
	fmt.Println("download complated.")
	return nil
}

// HTTPDownload ..
type HTTPDownload struct {
	*template
}

func (h *HTTPDownload) download() {
	fmt.Println("http download")
}

func (h *HTTPDownload) save() {
	fmt.Println("http save")
}

// NewHTTPDownload ..
func NewHTTPDownload() Downloader {
	httpDownload := new(HTTPDownload)
	tpl := new(template)
	tpl.implement = httpDownload
	httpDownload.template = tpl
	return httpDownload
}

type FTPDownload struct {
	*template
}

func (f *FTPDownload) download() {
	fmt.Println("ftp download")
}

func (f *FTPDownload) save() {
	fmt.Println("ftp save")
}

// NewFTPDownload ...
func NewFTPDownload() Downloader {
	ftpDonload := new(FTPDownload)
	tpl := new(template)
	tpl.implement = ftpDonload
	ftpDonload.template = tpl
	return ftpDonload
}
