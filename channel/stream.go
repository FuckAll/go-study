package channel

func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for _, value := range values {
			select {
			case <-done:
				return
			case ch <- value:
			}
		}
	}()
	return ch
}

func takeN(done <-chan struct{}, in <-chan interface{}, n int) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		var i int
		for {
			select {
			case <-done:
				return
			case a, ok := <-in:
				if !ok {
					return
				}
				if i < n {
					ch <- a
					i++
				}
			}
		}
	}()
	return ch
}

func takeFn(done <-chan struct{}, in <-chan interface{}, f func(v interface{}) bool) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case a, ok := <-in:
				if !ok {
					return
				}
				if f(a) {
					ch <- a
				}
			}
		}
	}()
	return ch
}

func takeWhile(done <-chan struct{}, in <-chan interface{}, f func(v interface{}) bool) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case a, ok := <-in:
				if !ok {
					return
				}
				if !f(a) {
					return
				}
				ch <- a
			}
		}
	}()
	return ch
}

func skipN(done <-chan struct{}, in <-chan interface{}, n int) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		var i int
		for {
			select {
			case <-done:
				return
			case a, ok := <-in:
				if !ok {
					return
				}
				i++
				if i <= n {
					continue
				}
				ch <- a
			}
		}
	}()
	return ch
}

func skipFn(done <-chan struct{}, in <-chan interface{}, f func(v interface{}) bool) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case a, ok := <-in:
				if !ok {
					return
				}
				if f(a) {
					continue
				}
				ch <- a
			}
		}
	}()
	return ch
}

func skipWhile(done <-chan struct{}, in <-chan interface{}, f func(v interface{}) bool) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case a, ok := <-in:
				if !ok {
					return
				}
				if f(a) {
					continue
				}
				ch <- a
			}
		}
	}()
	return ch
}
