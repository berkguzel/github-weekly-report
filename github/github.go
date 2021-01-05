package github

import (
	"fmt"
	"context"
	"time"
	"strings"

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
var repos []string
var initialRepo = make(map[string]*Repository)
var observerRepo = make(map[string]*Repository)

func InitialRepository(name string) *Repository{

	initRepo := &Repository{}
	i := Authentication(initRepo, name)
	initialRepo[name] = i

	return i
}

func ObserverRepository(name string) *Repository{
	
	obsRepo := &Repository{}
	o :=Authentication(obsRepo, name)
	observerRepo[name] = o

	return o
}

func RepositoryArray(repository string) []string {
	repo := strings.Split(repository, ",")
	for _, element := range repo {
		repos = append(repos, element)
	}

	return repos
}
func Authentication(r *Repository, repox string) *Repository {
	Time := time.Now()
	user := ParseArgs()
	token, _ := user["token"]
	owner, _ := user["owner"]
	repository :=  repox


	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	resp, _, err := client.Repositories.Get(ctx, owner, repository)
	if err != nil{
		fmt.Print(err)
	}
	
	repo := Repository{
		Name : *resp.Name,
		ForksCount :*resp.ForksCount,
		OpenIssuesCount : *resp.OpenIssuesCount,
		StargazersCount : *resp.StargazersCount,
		Time : Time.Format("2006.01.02 15:04:05"),
	}
	return &repo
	
}
