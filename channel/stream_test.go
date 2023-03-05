package channel

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_asStream(t *testing.T) {
	Convey("testAsStream", t, func() {
		done := make(chan struct{})
		testData := []interface{}{1, 2, 3, 4}
		got := asStream(done, testData...)
		var result []interface{}
		for i := 0; i < len(testData); i++ {
			result = append(result, <-got)
		}
		So(result, ShouldResemble, testData)
		t.Logf("result %+v", result)
	})
}

func Test_takeN(t *testing.T) {
	Convey("test task N", t, func() {
		done := make(chan struct{})
		testData := []interface{}{1, 2, 3, 4}
		got := asStream(done, testData...)
		rCh := takeN(done, got, 3)
		var result []interface{}
		for v := range rCh {
			result = append(result, v)
		}
		So(result, ShouldResemble, []interface{}{1, 2, 3})
	})
}

func Test_takeFn(t *testing.T) {
	Convey("test task Fn", t, func() {
		done := make(chan struct{})
		testData := []interface{}{1, 2, 3, 4}
		stream := asStream(done, testData...)
		rCh := takeFn(done, stream, func(v interface{}) bool {
			if i, ok := v.(int); ok {
				switch i {
				case 1, 4:
					return true
				}
			}
			return false
		})
		var result []interface{}
		for v := range rCh {
			result = append(result, v)
		}
		So(result, ShouldResemble, []interface{}{1, 4})
	})
}

func Test_takeWhile(t *testing.T) {
	Convey("test task Fn", t, func() {
		done := make(chan struct{})
		testData := []interface{}{1, 2, 3, 4}
		stream := asStream(done, testData...)
		rCh := takeWhile(done, stream, func(v interface{}) bool {
			if i, ok := v.(int); ok && i < 2 {
				return true
			}
			return false
		})
		var result []interface{}
		for v := range rCh {
			result = append(result, v)
		}
		So(result, ShouldResemble, []interface{}{1})
	})
}

func Test_skipN(t *testing.T) {
	Convey("test skipN", t, func() {
		done := make(chan struct{})
		testData := []interface{}{1, 2, 3, 4}
		stream := asStream(done, testData...)
		rCh := skipN(done, stream, 2)
		var result []interface{}
		for v := range rCh {
			result = append(result, v)
		}
		So(result, ShouldResemble, []interface{}{3, 4})
	})
}

func Test_skipFn(t *testing.T) {
	Convey("test skip fn", t, func() {
		done := make(chan struct{})
		testData := []interface{}{1, 2, 3, 4}
		stream := asStream(done, testData...)
		rCh := skipFn(done, stream, func(v interface{}) bool {
			if i, ok := v.(int); ok && i <= 2 {
				return true
			}
			return false
		})
		var result []interface{}
		for v := range rCh {
			result = append(result, v)
		}
		So(result, ShouldResemble, []interface{}{3, 4})
	})
}

func Test_skipWhile(t *testing.T) {
	Convey("test skip while", t, func() {
		done := make(chan struct{})
		testData := []interface{}{1, 2, 3, 4}
		stream := asStream(done, testData...)
		rCh := skipFn(done, stream, func(v interface{}) bool {
			if i, ok := v.(int); ok && i <= 2 {
				return true
			}
			return false
		})
		var result []interface{}
		for v := range rCh {
			result = append(result, v)
		}
		So(result, ShouldResemble, []interface{}{3, 4})
	})
}
