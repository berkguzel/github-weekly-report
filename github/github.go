package github

import (
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

}
var repos []string
var initialRepo []Repository
var observerRepo []Repository

// InitialRepository() runs starting of the time interval to be
// a referance to make comparison
func InitialRepository(sizeOfRepos int, arrayofRepos []string) ([]Repository, error) {

	initRepo := &Repository{}

	initialRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		v, _ := initRepo.GetValues(name)

		initialRepo = append(initialRepo, Repository{
			Name: v.Name,
			ForksCount: v.ForksCount,
			OpenIssuesCount: v.OpenIssuesCount,
			StargazersCount: v.StargazersCount,
			Time: v.Time,
			Fork: v.Fork,
		})
	}
	return initialRepo, nil
}

// ObserverRepository() runs when notification time has come
func ObserverRepository(sizeOfRepos int, arrayofRepos []string) ([]Repository, error){

	obsRepo := &Repository{}

	observerRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		v, err := obsRepo.GetValues(name)
		if err != nil {
			return nil, err
		}
		observerRepo = append(observerRepo, Repository{
			Name: v.Name,
			ForksCount: v.ForksCount,
			OpenIssuesCount: v.OpenIssuesCount,
			StargazersCount: v.StargazersCount,
			Time: v.Time,
			Fork: v.Fork,
		})
	}

	return observerRepo, nil
}

// RepositoryArray() appends name of the repositories to repos array
func RepositoryArray(repository string) ([]string, error) {
	
	if repository == "all" || repository == "ALL" {
		v, _ := GetAllRepositories()
		return v, nil
	}
	repo := strings.Split(repository, ",")
	for _, element := range repo {
		if element != "" {
			repos = append(repos, element)
		}
	}

	return repos, nil
}

// GetAllRepositories() runs when all option selected for repository
func GetAllRepositories() ([]string, error) {

	_,_, fork := Flags()
	client, ctx, owner := Authentication() 

	resp, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}

	for _, repo := range resp{
		if *repo.Owner.Login == owner &&
		fork == *repo.Fork {
			repos = append(repos, *repo.Name)
		}
	}

	return repos, nil
}


// picks up the values on your repositories
func (r *Repository) GetValues(repox string)  (*Repository, error){

	client, ctx, owner := Authentication()
	resp, _, err := client.Repositories.Get(ctx, owner, repox)
	if err != nil{
		return nil, err
	}

	r.Name = *resp.Name
	r.ForksCount = *resp.ForksCount
	r.OpenIssuesCount = *resp.OpenIssuesCount
	r.StargazersCount = *resp.StargazersCount
	r.Fork = *resp.Fork
	
	return r, nil
}

// Authentication() creates a connection between your Github account
func Authentication() (*github.Client, context.Context, string) {

	user := ParseArgs()
	token, _ := user["token"]
	owner, _ := user["owner"]

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client, ctx, owner
	
}
