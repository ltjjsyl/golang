package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	start1 := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Printf("method1: %v time=%d\n", s, time.Since(start1).Nanoseconds())

	start2 := time.Now()
	fmt.Printf("method2: %v time=%d\n", strings.Join(os.Args[0:], " "), time.Since(start2).Nanoseconds())
}
