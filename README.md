# ulimit
ulimit.SetMax() is go code that allows users to set the maximum possible filesystem ulimit.

In theory, when we run the following code, we should be able to set the ulimit (rLimit.Cur) to any number up to rLimit.Max: 

```
var rLimit syscall.Rlimit
err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
...
```

Unfortunately, if we try to set rLimit.Cur as rLimit.Max, the system may return an error if the actual limit is lower than rLimit.Max.

Therefore ulimit.SetMax() uses brute force to find and then set the actual filesystem ulimit.
