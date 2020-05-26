package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T){

	var x interface{} = 1
	var y interface{} = 2
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		t.Error("type not equal.")
	}
}

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func TestNewError(t *testing.T) {
	var rpcErr error = NewRPCError(400, "unknown err") // typecheck1
	err := AsErr(rpcErr) // typecheck2
	println(err )
}

func NewRPCError(code int64, msg string) error {
	var xx *RPCError = nil
	x1 := reflect.TypeOf(xx)
	x2 := reflect.TypeOf(&RPCError{})
	_, _ = x1, x2
	ei := (interface{})(nil)
	fmt.Printf("reflect nil : %v \n", reflect.TypeOf(nil) == reflect.TypeOf(ei) )
	fmt.Printf("ei: %v \n ", ei == (*error)(nil) )
	println(reflect.TypeOf(&RPCError{}).Implements(reflect.TypeOf((*error)(nil)).Elem()))
	return xx
	//return &RPCError{ // typecheck3
	//	Code:    code,
	//	Message: msg,
	//}
}

func AsErr(err error) error {
	return err
}