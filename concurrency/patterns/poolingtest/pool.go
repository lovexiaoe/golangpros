package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("池已被关闭")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("池大小不能小于0.")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("获取:", "共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	default:
		log.Println("获取:", "新资源")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("向资源列表中写入")
	default:
		log.Println("资源列表满了，关闭该资源")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
