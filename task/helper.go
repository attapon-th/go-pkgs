package task

import (
	"fmt"
	"log"
)

func panicHendler(t *Tasks) func(any) {
	return func(x any) {
		if x != nil {
			err := fmt.Errorf("Task Panic Error: %v", x)
			log.Fatal(err)
		}
	}
}
