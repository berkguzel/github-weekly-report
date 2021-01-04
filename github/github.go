package github

import (
	"fmt"
	"context"
	"time"

	"golang.org/x/oauth2"
	"github.com/google/go-github/github" 
)

type Repository struct{
	Name string
	ForksCount int
	OpenIssuesCount int
	StargazersCount int
	Time string

}

func InitialRepository() *Repository{
	
	initRepo := &Repository{}
	r := Authentication(initRepo)
	return r
}

func ObserverRepository() *Repository{
	
	obsRepo := &Repository{}
	o :=Authentication(obsRepo)
	return o
}


func Authentication(r *Repository) *Repository {
	Time := time.Now()
	user := parseArgs()
	token, _ := user["token"]
	owner, _ := user["owner"]
	repository, _ := user["repository"]


	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.Get(ctx, owner, repository)
	if err != nil{
		fmt.Print(err)
	}
	
	repo := Repository{
		Name : *repos.Name,
		ForksCount :*repos.ForksCount,
		OpenIssuesCount : *repos.OpenIssuesCount,
		StargazersCount : *repos.StargazersCount,
		Time : Time.Format("2006.01.02 15:04:05"),
	}
	return &repo
	
}
