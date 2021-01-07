package github

import (
	"os"
	"log"
)
var (
	owner string
	repository string
	token string
	fork bool
)

func ParseArgs() (map[string]string, bool) {
	//TODO get interval day as arg
	 
	user := make(map[string]string)
	
	token = os.Getenv("ACCESS_TOKEN")
	if token == ""{
		log.Fatal("Github access token is not set")
	}
	user["token"] = token
	
	owner = os.Getenv("OWNER")
	if owner == ""{
		log.Fatal("Owner of the repositories is not set")
	}
	user["owner"] = owner
	
	forked := os.Getenv("FORKED")
	if forked == "" {
		log.Fatal("Forked option is not set")
	}else if forked == "false"{
		fork = false
	}else{
		fork = true
	}
	
	repository = os.Getenv("REPOSITORY")
	if repository == ""{
		log.Fatal("Repository is not set")
	}
	user["repository"] = repository
	// TODO be sure user seperated repositories with comma ,

	return user, fork
}