package main

import (
	"flag"
	"fmt"
)

func main() {
	count := flag.Int("count", 1, "count is int")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("need argument")
		return
	}

	for i := 0; i < *count; i++ {
		fmt.Println(flag.Arg(0))
	}
}
