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
var initialRepo []*Repository 
var observerRepo []*Repository

// InitialRepository() runs starting of the time interval to be
// a referance to make comparison
func InitialRepository(sizeOfRepos int, arrayofRepos []string) []*Repository {

	initRepo := &Repository{}

	initialRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		initialRepo = append(initialRepo, initRepo.Authentication(name))
	}

	return initialRepo
}

// ObserverRepository() runs when notification time has come
func ObserverRepository(sizeOfRepos int, arrayofRepos []string) []*Repository{

	obsRepo := &Repository{}

	observerRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		observerRepo = append(observerRepo,obsRepo.Authentication(name))
	}

	return observerRepo
}

// RepositoryArray() appends name of the repositories to repos array
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

// GetAllRepositories() runs when all option selected for repository
// We use repository name to call Authentication() because
// it must be call with a repository name
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

// Authentication() creates a connection between your Github account
// picks up the values on your repositories
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

	return r
	
}
