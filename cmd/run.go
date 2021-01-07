package main

import (
	"fmt"
	"time"

	"github-weekly-report/github"
)
//TODO notify me after 5 minutes to test system
//TODO notify me if there is difference
//TODO notify me only star or star and issue etc.

var initialRepo []*github.Repository 
var observerRepo []*github.Repository
func RunOnce(sizeOfRepos int, arrayofRepos []string) []*github.Repository{

	initialRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		//fmt.Println(github.InitialRepository(name))
		initialRepo = append(initialRepo, github.InitialRepository(name))
	}
	return initialRepo
}
func RunPeroidically(sizeOfRepos int, arrayofRepos []string) []*github.Repository{

	observerRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		//fmt.Println(github.ObserverRepository(name))
		observerRepo = append(observerRepo, github.ObserverRepository(name))
	}	
	return observerRepo
}
func WeekDay() time.Weekday{

	return time.Now().Weekday()
}

func Diff(i []*github.Repository, o []*github.Repository ){
	// TODO get initial and observer structes and compare
	for r , v := range o {
		star := v.StargazersCount - i[r].StargazersCount
		forks := v.ForksCount - i[r].ForksCount
		issue := v.OpenIssuesCount - i[r].OpenIssuesCount
		fmt.Println(v.Name, star, forks, issue)
	}
}