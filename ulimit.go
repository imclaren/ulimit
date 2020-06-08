package ulimit

import (
	"fmt"
	"syscall"
)

// SetMax sets the ulimit (rLimit.Cur) to i if the current ulimit is less than i
// If i is zero, the ulimit (rLimit.Cur) is set to the maximum avalable rLimit
// macOS maximum seems to be 10240
func SetMax(i uint64) error {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return fmt.Errorf("Error Getting Rlimit: %s", err.Error())
	}
	//rLimit.Max = 999999
	//rLimit.Cur = 999999

	// set max limit if i == 0
	if i == 0 {
		i = rLimit.Max
		lowerBound := uint64(0)
		for {
			rLimit.Cur = i
			err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
			if err == nil {
				if lowerBound == 0 {
					lowerBound = i
					i = i*10
				} else {
					break
				}
			}
			if lowerBound == 0 {
				i = i/10
			} else {
				i--
			}
			if i <= 1 {
				return fmt.Errorf("Error Setting Rlimit: %s", err.Error())
			}
		}
	}

	if rLimit.Cur < i {
		rLimit.Cur = i
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			return fmt.Errorf("Error Setting Rlimit: %s", err.Error())
		}
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return fmt.Errorf("Error Getting Rlimit: %s", err.Error())
	}
	if rLimit.Cur < i {
		return fmt.Errorf("rLimit.Cur update failure - expected at least: %d result: %d", i, rLimit.Cur)
	}
	return nil
}
