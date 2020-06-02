package template

/*
	模板模式：一个抽象类公开定义了执行它的方法的方式/模板。
	实现一个下载类：下载类要有download()方法和save()方法。
*/
import "fmt"

// Downloader ...
type Downloader interface {
	Download() error
}

type DownloadImpl interface {
	download()
	save()
}

type DownloadTemplate struct {
	DownloadImpl // 使用匿名可以达到继承的效果
	uri          string
}

func (dt *DownloadTemplate) Download() error {
	dt.download()
	dt.save()
	return nil
}

// HTTPDownload ..
type HTTPDownload struct {
	*DownloadTemplate
}

func (h *HTTPDownload) download() {
	fmt.Println("http download")
}

func (h *HTTPDownload) save() {
	fmt.Println("http save")
}

// NewHTTPDownload ..
func NewHTTPDownload() Downloader {
	// 实际实现的方法是HTTPDownload
	httpDownload := new(HTTPDownload)
	tpl := new(DownloadTemplate)
	tpl.DownloadImpl = httpDownload
	httpDownload.DownloadTemplate = tpl
	return httpDownload
}

// FTPDownload ...
type FTPDownload struct {
	*DownloadTemplate
}

func (f *FTPDownload) download() {
	fmt.Println("ftp download")
}

func (f *FTPDownload) save() {
	fmt.Println("ftp save")
}

// NewFTPDownload ...
func NewFTPDownload() Downloader {
	ftpDownload := new(FTPDownload)
	tpl := new(DownloadTemplate)
	tpl.DownloadImpl = ftpDownload
	ftpDownload.DownloadTemplate = tpl
	return ftpDownload
}
