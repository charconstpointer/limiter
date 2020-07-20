package limiter

import (
	"context"
)

//Limiter helps you to limit count of go routines used to execute given task
type Limiter struct {
	c chan struct{}
}

//NewLimiter creates initial job pool, and returns new Limiter
func NewLimiter(limit int) *Limiter {
	limiter := &Limiter{c: make(chan struct{}, limit)}
	for i := 0; i < limit; i++ {
		limiter.c <- struct{}{}
	}
	return limiter
}

//Run executes job according to provided rate limiting settings
func (l *Limiter) Run(f func()) {
	<-l.c
	go func() {
		defer func() {
			l.done()
		}()
		f()
	}()
}

//Wait waits until context is cancelled
func (l *Limiter) Wait(ctx context.Context) {
	<-ctx.Done()
}

func (l *Limiter) done() {
	l.c <- struct{}{}
}
