package main

import (
	"fmt"
	services "./services"
	"./structs"
	"regexp"
	"log"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	LoadEnv()
	owner := os.Getenv("OWNER")
	repos_pointer := os.Getenv("REPOS")	
	repos := parseRepos(repos_pointer)
	var branch_list [] structs.Branch

	for _, repo := range repos {
		branch := services.ParsePayload(owner, repo)
		fmt.Printf("%+v",branch)
		branch_list = append(branch_list, branch)
	}

	services.CreateAndWriteCsv(branch_list)
}

func parseRepos(repos string) []string{
	params := regexp.MustCompile(", ")
	result_repos := params.Split(repos, -1)

	return result_repos
}

func LoadEnv() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}