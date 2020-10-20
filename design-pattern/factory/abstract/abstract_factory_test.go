package abstract

import (
	"testing"
)

func Test_createFactory(t *testing.T) {
	factory := createFactory("good")
	html := factory.createHtml("#Hello\nHello, world!")
	html.save("fast.html")

	word := factory.createWord("#Hello\nHello, world!")
	word.save("fast.word")

	factory = createFactory("fast")
	html = factory.createHtml("#Hello\nHello, world!")
	html.save("fast.html")

	word = factory.createWord("#Hello\nHello, world!")
	word.save("fast.word")
}
