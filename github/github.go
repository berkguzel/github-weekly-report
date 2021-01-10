package github

import (
	"log"
	"context"
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
	Fork bool
	WatchersCount int

}
var repos []string
func InitialRepository(name string) *Repository {

	initRepo := &Repository{}
	i := initRepo.Authentication(name)

	return i
}

func ObserverRepository(name string) *Repository {

	obsRepo := &Repository{}
	o := obsRepo.Authentication(name)

	return o
}

func RepositoryArray(repository string) []string {
	
	if repository == "all" || repository == "ALL" {
		return GetAllRepositories()
	}
	repo := strings.Split(repository, ",")
	for _, element := range repo {
		if element != "" {
			repos = append(repos, element)
		}
	}

	return repos
}

func GetAllRepositories() []string {

	user, fork := ParseArgs()
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
		log.Fatal(err)
	}

	for _, repo := range resp{
		if *repo.Owner.Login == owner &&
		fork == *repo.Fork {
			repos = append(repos, *repo.Name)
		}
	}

	return repos
}
func (r *Repository) Authentication(repox string) *Repository {

	user, _ := ParseArgs()
	token, _ := user["token"]
	owner, _ := user["owner"]
	repository :=  repox

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	resp, _, err := client.Repositories.Get(ctx, owner, repository)
	if err != nil{
		log.Fatal(err)
	}

	r.Name = *resp.Name
	r.ForksCount = *resp.ForksCount
	r.OpenIssuesCount = *resp.OpenIssuesCount
	r.StargazersCount = *resp.StargazersCount
	r.Fork = *resp.Fork
	r.WatchersCount = *resp.WatchersCount
	
	return r
	
}
