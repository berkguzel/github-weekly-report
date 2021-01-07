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
	
	_, fork := github.ParseArgs()
	initialRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		r :=  github.InitialRepository(name)
		if fork == r.Fork{
			initialRepo = append(initialRepo, r)
		}
	}

	return initialRepo
}
func RunPeroidically(sizeOfRepos int, arrayofRepos []string) []*github.Repository{

	_, fork := github.ParseArgs()
	observerRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		r := github.ObserverRepository(name)
		if fork == r.Fork {
			observerRepo = append(observerRepo, r)
		}
	}	

	return observerRepo
}
func WeekDay() time.Weekday{

	return time.Now().Weekday()
}

func Diff(i []*github.Repository, o []*github.Repository ){

	for r , v := range o {
		star := v.StargazersCount - i[r].StargazersCount
		forks := v.ForksCount - i[r].ForksCount
		issue := v.OpenIssuesCount - i[r].OpenIssuesCount
		fork := v.Fork
		fmt.Println(v.Name, star, forks, issue, fork)
	}
}