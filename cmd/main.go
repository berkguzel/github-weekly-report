package main

import (
	"time"

	"github-weekly-report/github"
)


func main() {

	arg, _ := github.ParseArgs()
	repository := arg["repository"]
	sizeOfRepos := len(github.RepositoryArray(repository))
	arrayofRepos := github.RepositoryArray(repository)
	
	initialRepo = RunOnce(sizeOfRepos, arrayofRepos)
	for c := time.Tick(20 * time.Second); ; <-c { 
		observerRepo = RunPeroidically(sizeOfRepos, arrayofRepos)
		Diff(initialRepo, observerRepo)
	}
}