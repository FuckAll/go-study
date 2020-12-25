package log

import (
	"log"
	"testing"
)

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func Test_Log(t *testing.T) {
	log.Println("good")

	log.Fatalln("wonderful")

	log.Println("hah")
}
