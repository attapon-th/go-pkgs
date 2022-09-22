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
    "log"
    "time"
    "github.com/attapon-th/go-pkgs/task"
)

func main() {
    t1 := runTasksSample("     wait task", 2, 5)
    t2 := runTasksSample("backgroud task", 2, 10) 
    defer t1.Close() // close 
    defer t2.Close()

    go t2.RunTasks() // call run task with goroutine(backgroup process)

    t1.RuRunTasksAndEnd() // wait task t1  complated

    // add worker increment 2 summary = 4
    t2.SetWorker(4)
    t2.Wait() // wait backgroup task complated

}

// runTasksSample - Example How to Create Task
//
//  @return *task.Tasks
func runTasksSample(name string, worker, maxTask int) *task.Tasks {
    t, err := task.NewTasks(worker)
    if err != nil {
        log.Fatalln(err)
    }
    for i := 0; i < maxTask; i++ {
        prefix := fmt.Sprintf("Task[%s] %d of %d", name, i+1, maxTask)
        t.AddTask(funcPrintLoop(prefix, 5))

    }
    t.AddTask(t.Close)
    return t
}

// funcPrintLoop - Example function to AddTask
//
//  @param prefix
//  @param max
//  @return func() **recommend return value of func() if function is dynamic parse parameter
func funcPrintLoop(prefix string, max int) func() {
    return func() {
        fmt.Println(prefix, " - is running...")
        for i := 0; i < max; i++ {
            time.Sleep(time.Millisecond * 500)
        }
        fmt.Println(prefix, " - is completed")
    }
}
```
