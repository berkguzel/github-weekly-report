package main

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github-weekly-report/github"
)
//TODO notify me after 5 minutes to test system
//TODO notify me if there is difference
//TODO notify me only star or star and issue etc.
//TODO notify total star, issuei fork 
var initialRepo []*github.Repository 
var observerRepo []*github.Repository

func RunOnce(sizeOfRepos int, arrayofRepos []string) []*github.Repository{
	
	initialRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		initialRepo = append(initialRepo, github.InitialRepository(name))
	}

	return initialRepo
}
func RunPeroidically(sizeOfRepos int, arrayofRepos []string) []*github.Repository{

	observerRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		observerRepo = append(observerRepo, github.ObserverRepository(name))
	}

	return observerRepo
}

func Diff(i []*github.Repository, o []*github.Repository ){

	arg, _ := github.ParseArgs()
	bot, err := tgbotapi.NewBotAPI(arg["tgBotToken"])
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	chatId, _ := strconv.ParseInt(arg["chatId"], 10, 64)

	var message string
	for r , v := range o {
		name := v.Name
		star := v.StargazersCount - i[r].StargazersCount
		forks := v.ForksCount - i[r].ForksCount
		issue := v.OpenIssuesCount - i[r].OpenIssuesCount
		watchers := v.WatchersCount - i[r].WatchersCount
		message = message + "\n" + fmt.Sprintf("Repository: %s \n New Stars: %d \n New Forks: %d \n New Issues: %d \n New Watchers: %d \n", name, star, forks, issue, watchers )
	}
	message = fmt.Sprintf("Hi, %s! Here is what happened in your repositories. \n %s ", arg["owner"], message)
	msg := tgbotapi.NewMessage(chatId, message)
	bot.Send(msg)
	
}

func Filter(){
	//TODO filter the results
}