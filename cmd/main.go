package main

import (
	"time"
	
	"github-weekly-report/github"
)

var (
	timeD time.Duration
	interval int
)
func main() {

	arg, _ := github.ParseArgs()
	repository := arg["repository"]
	sizeOfRepos := len(github.RepositoryArray(repository))
	arrayofRepos := github.RepositoryArray(repository)
	interval, timeD = github.Flags()
	
	initialRepo = RunOnce(sizeOfRepos, arrayofRepos)
	for c := time.Tick(time.Duration(interval) * timeD); ; <-c { 

		observerRepo = RunPeroidically(sizeOfRepos, arrayofRepos)
		Diff(initialRepo,observerRepo)
		initialRepo = RunOnce(sizeOfRepos, arrayofRepos)
	
	}
	
}
