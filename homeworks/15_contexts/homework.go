package homework15

import (
	"context"
	"sync"
)

type Group struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	once   sync.Once
	err    error
}

func NewErrGroup(ctx context.Context) (*Group, context.Context) {
	contex, cancel := context.WithCancel(ctx)
	return &Group{
		ctx:    contex,
		cancel: cancel,
	}, contex
}

func (g *Group) Go(action func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()
		if err := action(); err != nil {
			g.once.Do(func() {
				g.cancel()
				g.err = err
			})
		}

	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	return g.err
}
