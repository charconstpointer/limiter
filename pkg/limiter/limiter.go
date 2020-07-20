package limiter

//Limiter helps you to limit count of go routines used to execute given task
type Limiter struct {
	c chan int
}

//NewLimiter creates initial job pool, and returns new Limiter
func NewLimiter(limit int) *Limiter {
	limiter := &Limiter{c: make(chan int, limit)}
	for i := 0; i < limit; i++ {
		limiter.c <- i
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

func (l *Limiter) done() {
	l.c <- 1
}
