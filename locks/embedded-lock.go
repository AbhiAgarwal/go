package main

import (
	"fmt"
	"sync"
)

var hits struct {
	sync.Mutex
	n int
}

func main() {
	// Lock means this is using it and nothing else can
	// Unlock means let someone else us it
	hits.Lock()
	hits.n++
	hits.n++
	fmt.Println(hits.n)
}
