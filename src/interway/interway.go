package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"

	"reflect"

	"net/http"

	"encoding/json"

	"github.com/golang/protobuf/proto"
	gctx "github.com/gorilla/context"
	"github.com/pborman/uuid"
	"github.com/urfave/negroni"
	"github.com/wothing/log"
)

const svcName = "interway"
const port = 50032

const (
	printStack bool = false
	stackAll   bool = false
	stackSize  int  = 1024 * 8
)

func respondBytes(rw http.ResponseWriter, r *http.Request, data []byte) {
	tid := gctx.Get(r, "tid").(string)
	l := len(data)
	if l > 10000 {
		l = 10000
	}
	log.Tinfof(tid, "responding, response=%s", string(data[:l]))
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

func RecoveryMiddleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				stack := make([]byte, stackSize)

				stack = stack[:runtime.Stack(stack, stackAll)]
				log.Error("panic http request: %s, err=%s, stack:\n%s", r.RequestURI, err, string(stack))
				if printStack {
					fmt.Fprintf(rw, "PANIC: %s\n%s", err, stack)
				} else {
					fmt.Fprintf(rw, "internal error")
				}
			}
		}()
		next(rw, r)
	}
}

func respondObject(rw http.ResponseWriter, r *http.Request, obj interface{}) {
	message, err := json.Marshal(obj)
}
func RespondMessage(rw http.ResponseWriter, r *http.Request, message interface{}) {
	switch m := message.(type) {
	case proto.Message:
		fmt.Printf("wonderful, %v", m)

	case string:
		respondBytes(rw, r, []byte(m))
	default:
		respondBytes(rw, r, message)
	}
}

func LogMiddleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()
		tid := uuid.New()
		rw.Header().Set("X-Request-ID", tid)
		version := r.Header().Get("X-App-Version")
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Tinfo(tid, "error on reading request body, from=%v, method=%v, remote=%v, agent=%v, version=%s", r.RequestURI, r.Method, r.RemoteAddr, r.UserAgent(), version)

		}
	}

}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n := negroni.New()
	// n.Use(m)

	fmt.Println(reflect.TypeOf(n))

}
