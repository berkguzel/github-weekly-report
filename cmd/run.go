package main

import (
	"fmt"
	"time"

	"github-weekly-report/github"
)

func RunOnce(sizeOfRepos int, arrayofRepos []string) {

	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		fmt.Println(github.InitialRepository(name))
	}
}
func RunPeroidically(sizeOfRepos int, arrayofRepos []string) {

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
}