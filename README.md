# ulimit
ulimit.SetMax() allows users to set the maximum possible filesystem ulimit ([GoDoc](https://godoc.org/github.com/imclaren/ulimit)).

```
oldLimit, err := ulimit.Get()
if err != nil {
	log.Fatal(err)
}
err = SetMax()
if err != nil {
	log.Fatal(err)
}
newLimit, err := Get()
if err != nil {
	log.Fatal(err)
}
fmt.Println(oldLimit, newLimit)
// -> 256 24576
```

In theory, when we run the following go code, we should be able to set the ulimit (rLimit.Cur) to any number up to rLimit.Max: 

```
var rLimit syscall.Rlimit
err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
...
```

Unfortunately, if we try to set rLimit.Cur as rLimit.Max, the system may return an error because the actual limit is lower than rLimit.Max.  Therefore ulimit.SetMax() uses brute force to find and then set the maximum possible filesystem ulimit.