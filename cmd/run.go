package main

import (
	"fmt"
	"time"

	"github-weekly-report/github"
)

func RunOnce(){

	arg := github.ParseArgs()
	repository := arg["repository"]
	sizeOfRepos := len(github.RepositoryArray(repository))
	arrayofRepos := github.RepositoryArray(repository)
	
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		fmt.Println(github.InitialRepository(name))
	}

}
func RunPeroidically(){

	arg := github.ParseArgs()
	repository := arg["repository"]
	sizeOfRepos := len(github.RepositoryArray(repository))
	arrayofRepos := github.RepositoryArray(repository)
	
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		fmt.Println(github.ObserverRepository(name))
	}
	
}
func WeekDay() time.Weekday{

	return time.Now().Weekday()
}

func Diff(){

	// TODO get initial and observer structes and compare


	//type initRepo = github.InitialRepository()
	//type obsRepo = github.ObserverRepository()

}