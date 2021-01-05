package main

import (
	"fmt"
	"time"
)

func main() {
	
	RunOnce()
	for c := time.Tick(10 * time.Second); ; <-c { 
		RunPeroidically()
	}
	
}










