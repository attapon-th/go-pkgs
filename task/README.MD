# `task` -- Task parallel workflow

## Requirement
-  `go` version 1.18 or more

## Dependencies package

- [Ants](https://github.com/panjf2000/ants) - Library ants implements a goroutine pool with fixed capacity, managing and recycling a massive number of goroutines, allowing developers to limit the number of goroutines in your concurrent programs.

- [sync.WaitGroup](https://pkg.go.dev/sync) - Package sync provides basic synchronization primitives such as mutual exclusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.  
Values containing the types defined in this package should not be copied.

## How to install

## How to use

### Import Package
```go
import (
    "fmt"
    "log"
    "time"

    "github.com/attapon-th/go-pkgs/task"
)
```

### Create Task



## [Example](./_example/main.go)

```go
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/attapon-th/go-pkgs/task"
)

func main() {
	// runtime.
	runTasksSample()
}

func runTasksSample() {
	t1, _ := task.New(2) // call summit is task start running

	for i := 1; i <= 10; i++ {
		t1.Submit(funcPrintLoop("Summit Task "+strconv.Itoa(i), 5))
	}
	time.Sleep(time.Second)
	defer t1.Release() // close

	for i := 1; i <= 5; i++ {
		t1.Submit(funcPrintLoop("Other Task "+strconv.Itoa(i), 5))
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Wait...")
	for t1.GetCount() > 0 {
		fmt.Println(t1.GetCount())
		time.Sleep(time.Second)
	}
	t1.Wait()
}

// funcPrintLoop - Example function to AddTask
//
//	@return func() **recommend return value of func() if function is dynamic parse parameter
func funcPrintLoop(prefix string, waitSec int) func() {
	fmt.Println("Add ", prefix)
	return func() {
		// fmt.Println(prefix, " - is running...")
		for i := 0; i < waitSec; i++ {
			time.Sleep(time.Millisecond * 500)
		}
		fmt.Println(prefix, " - is completed")
	}
}

```
