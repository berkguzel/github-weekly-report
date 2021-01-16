package main

import (
	"time"
	
	"github-weekly-report/github"
)

var (
	timeD time.Duration
	interval int
	initialRepo []*github.Repository 
	observerRepo []*github.Repository
)
func main() {

	arg, _ := github.ParseArgs()
	repository := arg["repository"]
	sizeOfRepos := len(github.RepositoryArray(repository))
	arrayofRepos := github.RepositoryArray(repository)
	
	interval, timeD = github.Flags()
	if interval == 0 {
		interval = 7 * 24 
		timeD = time.Hour
	}

	
	initialRepo = github.InitialRepository(sizeOfRepos, arrayofRepos)
	for c := time.Tick(time.Duration(interval) * timeD); ; <-c { 

		observerRepo = github.ObserverRepository(sizeOfRepos, arrayofRepos)
		Diff(initialRepo,observerRepo)
		initialRepo = github.InitialRepository(sizeOfRepos, arrayofRepos)
	
	}
	
}
