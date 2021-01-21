package github

import (

	"testing"
)

func TestInitialRepository(t *testing.T) {
	got, err := InitialRepository(0, nil)
	if err != nil {
		t.Error(err)
	}
	if got != nil {
		t.Errorf("Expected empty return from repositories \n %v", err)	
	}
}

func TestObserverRepository(t *testing.T) {
	got, err := ObserverRepository(0, nil)
	if err != nil {
		t.Error(err)
	}
	if got != nil {
		t.Errorf("Expected empty return from repositories  \n %v", err)
	}

}

func TestGetAllRepositories(t *testing.T) {
	v, err := GetAllRepositories()
	if err != nil {
		t.Error(err)
	}
	if v == nil {
		t.Errorf("Expected to collect all repositories")
	}
}

func TestRepositoryArray(t *testing.T) {
	v, err := RepositoryArray("")
	if err != nil {
		t.Error(err)
	}
	if v == nil {
		t.Errorf("Expected to return repositories")
	}
}
