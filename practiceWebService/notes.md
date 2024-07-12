# func Notify

```go
func Notify(c chan<- os.Signal, sig ...os.Signal)
```

Notify causes package signal to relay incoming signals to c.
If no signals are provided, all incoming signals will be relayed to c.
Otherwise, just the provided signals will.

Package signal will not block sending to c: the caller must
ensure that c has sufficient buffer space to keep up with
the expected signal rate. For a channel used for notification
of just one signal value, a buffer of size 1 is sufficient.
