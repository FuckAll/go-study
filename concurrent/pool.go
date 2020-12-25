package concurrent

import (
	"errors"
	"io"
	"sync"
)

type Pool struct {
	resources chan io.Closer
	factory   func() (io.Closer, error)

	// protects closed
	mu     sync.Mutex
	closed bool
}

var ErrPollClosed = errors.New("poll has been closed")

func NewPool(f func() (io.Closer, error), size int) (*Pool, error) {
	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   f,
		closed:    false,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case p, ok := <-p.resources:
		if !ok {
			return nil, ErrPollClosed
		}
		return p, nil
	default:
		return p.factory()
	}
}

func (p *Pool) Release(i io.Closer) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		// pool closed but i not closed
		i.Close()
		return
	}

	select {
	// may be panic , because p.resources close
	case p.resources <- i:

	default:
		i.Close()
	}
}

func (p *Pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)

	for re := range p.resources {
		re.Close()
	}
}
