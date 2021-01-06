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

//TODO forked=false option

func InitialRepository(name string) *Repository {

	initRepo := &Repository{}
	i := Authentication(initRepo, name)

	return i
}

func ObserverRepository(name string) *Repository {

	obsRepo := &Repository{}
	o :=Authentication(obsRepo, name)
	//fmt.Println(observerRepo[name].StargazersCount)
	return o
}

func RepositoryArray(repository string) []string {
	
	if repository == "all" || repository == "ALL" {
		return GetAllRepositories()
	}
	repo := strings.Split(repository, ",")
	for _, element := range repo {
		if element != ""{
			repos = append(repos, element)
		}

	}

	return repos
}

func GetAllRepositories() []string {

	user := ParseArgs()
	owner, _ := user["owner"]
	token, _ := user["token"]
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	resp, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, repo := range resp{
		if *repo.Owner.Login == owner {
			repos = append(repos, *repo.Name)
		}
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
