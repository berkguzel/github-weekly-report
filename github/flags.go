package github

import (
	"os"
	"log"
)
var (
	owner string
	repository string
)

func ParseArgs()map[string]string{

	//TODO get interval day as arg

	//TODO get all repositories of the profile 
	 
	user := make(map[string]string)
	
	token := os.Getenv("ACCESS_TOKEN")
	if token == ""{
		log.Fatal("Github access token is not set")
	}
	user["token"] = token
	
	owner := os.Getenv("OWNER")
	if owner == ""{
		log.Fatal("Owner of the repositories is not set")
	}
	user["owner"] = owner

	repository := os.Getenv("REPOSITORY")
	if repository == ""{
		log.Fatal("Repository is not set")
	}
	user["repository"] = repository
	// TODO be sure user seperated repositories with comma ,

	return user
}