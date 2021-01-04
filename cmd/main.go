package main

import (
	"github-weekly-report/github"
	"fmt"
	"time"

)

func main() {
	initRepo := github.InitialRepository()

	for c := time.Tick(1 * time.Second); ; <-c { 
		obsRepo := github.ObserverRepository()

	}
  
}










