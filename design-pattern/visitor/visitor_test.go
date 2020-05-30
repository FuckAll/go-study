package visitor

import "testing"

func TestVisitor(t *testing.T) {
	extractor := new(Extractor)

	resourceList := []ResourceFile{new(WordFile), new(PDFFile), new(PPTFile)}
	for _, r := range resourceList {
		r.accept(extractor)
	}

	compressor := new(Compressor)
	for _, r := range resourceList {
		r.accept(compressor)
	}
}
