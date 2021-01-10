package main

import (
	"time"
	
	"github-weekly-report/github"
)

//TODO ask multiple message or just one
// TODO message time, day

func main() {

	arg, _ := github.ParseArgs()
	repository := arg["repository"]
	sizeOfRepos := len(github.RepositoryArray(repository))
	arrayofRepos := github.RepositoryArray(repository)
	
	initialRepo = RunOnce(sizeOfRepos, arrayofRepos)
	for c := time.Tick(24 * 7 * time.Hour); ; <-c { 

		observerRepo = RunPeroidically(sizeOfRepos, arrayofRepos)
		Diff(initialRepo,observerRepo)
		initialRepo = RunOnce(sizeOfRepos, arrayofRepos)
	}
}
