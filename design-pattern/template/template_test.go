// 模板方法模式
package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	httpDownloader := NewHTTPDownload()
	httpDownloader.Download()

	fmt.Println("-----------------------------------")
	ftpDownloader := NewFTPDownload()
	ftpDownloader.Download()
}
