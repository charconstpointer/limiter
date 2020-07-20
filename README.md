# limiter
```
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	l := limiter.NewLimiter(2)
	for {
		l.Run(func() {
			fmt.Println("hi")
			cancel()
		})
	}
	l.Wait(ctx)
```
