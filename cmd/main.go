package main

import (
	"log"
	"time"

	"github.com/berkguzel/github-weekly-report/github"
)

var (
	initialRepo  []github.Repository
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

	notifyInterval, timeD, _ := github.Flags()
	if notifyInterval == 0 {
		notifyInterval = 7 * 24
		timeD = time.Hour
	}

	initialRepo, err = github.InitialRepository(sizeOfRepos, arrayofRepos)
	if err != nil {
		log.Fatal(err)
	}

	for c := time.Tick(time.Duration(notifyInterval) * timeD); ; <-c {

		observerRepo, err = github.ObserverRepository(sizeOfRepos, arrayofRepos)
		if err != nil {
			log.Fatal(err)
		}

		Diff(initialRepo, observerRepo)

		initialRepo, err = github.InitialRepository(sizeOfRepos, arrayofRepos)
		if err != nil {
			log.Fatal(err)
		}

	}

}
