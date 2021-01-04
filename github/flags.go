package github

import (
	"os"
	"log"

)
var (
	owner string
	repository string
)

func parseArgs()map[string]string{

	//TODO get interval day as arg

	//TODO get multiple repositores

	//TODO get all repositories of the profile 
	 
	user := make(map[string]string)
	
	token := os.Getenv("ACCESS_TOKEN")
	if token == ""{
		log.Fatal("Github access token is not set")
	}
	user["token"] = token
	
	owner := os.Getenv("OWNER")
	if owner != ""{
		user["owner"] = owner
	}
	repository := os.Getenv("REPOSITORY")
	/*
	fmt.Print(repository)
	rep := strings.Split(repository, " ")
	for _, element := range rep {
		fmt.Printf(element)
	}
	*/
	if repository != ""{
		user["repository"] = repository
	}


	return user
}