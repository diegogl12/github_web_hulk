package main

import (
	"fmt"
	services "./services"
	"./structs"
	"flag"
	"regexp"
)

func main() {
	owner := flag.String("owner", "diegogl12", "foo")
	repos_pointer := flag.String("repos", "", "foo")
	flag.Parse()
	repos := parseRepos(repos_pointer)
	var branch_list [] structs.Branch
	
	for _, repo := range repos {
		branch := services.ParsePayload(*owner, repo)
		fmt.Printf("%+v",branch)
		branch_list = append(branch_list, branch)
	}
	services.CreateAndWriteCsv(branch_list)
}

func parseRepos(repos *string) []string{
	params := regexp.MustCompile(",")
	result_repos := params.Split(*repos, -1)

	return result_repos
}

