package main

import (
	"time"
	"fmt"

	"github-weekly-report/github"
)

func main() {
	arg := github.ParseArgs()
	repository := arg["repository"]
	sizeOfRepos := len(github.RepositoryArray(repository))
	arrayofRepos := github.RepositoryArray(repository)

	
	RunOnce(sizeOfRepos, arrayofRepos)
	fmt.Println("*********************")
	for c := time.Tick(10 * time.Second); ; <-c { 
		RunPeroidically(sizeOfRepos, arrayofRepos)
	}

}