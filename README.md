# ulimit
ulimit allows users to run SetMax() to set the maximum filesystem ulimit.

In theory, when we run the following code, we should be able to set the ulimit (rLimit.Cur) to any number up to rLimit.Max: 

```
var rLimit syscall.Rlimit
err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
...
```

Unfortunately, if we try to set rLimit.Cur as rLimit.Max, the system can return an error if the actual limit is lower than rLimit.Max.

Therefore SetMax() uses brute force to find and then set the actual filesystem ulimit. 