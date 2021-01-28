package main

import (
	"time"
	"log"

	"github-weekly-report/github"
)

var (
	timeD time.Duration
	interval int
	initialRepo []github.Repository 
	observerRepo []github.Repository
)
func main() {

	arg := github.ParseArgs()
	repository := arg["repository"]
	arrayofRepos, err := github.RepositoryArray(repository)
	if err != nil {
		log.Fatal(err)
	}
	sizeOfRepos := len(arrayofRepos)

	interval, timeD, _ = github.Flags()
	if interval == 0 {
		interval = 7 * 24 
		timeD = time.Hour
	}

	initialRepo, err = github.InitialRepository(sizeOfRepos, arrayofRepos)
	if err != nil{
		log.Fatal(err)
	}

	for c := time.Tick(time.Duration(interval) * timeD); ; <-c { 

		observerRepo, err = github.ObserverRepository(sizeOfRepos, arrayofRepos)
		if err != nil {
			log.Fatal(err)
		}

		Diff(initialRepo,observerRepo)

		
		initialRepo, err = github.InitialRepository(sizeOfRepos, arrayofRepos)
		if err != nil {
			log.Fatal(err)
		}
	
	}
	
}
