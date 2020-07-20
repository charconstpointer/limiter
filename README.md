# limiter
```
	l := limiter.NewLimiter(2)
	for {
		l.Run(func() {
			fmt.Println("hi")
		})
	}
```
