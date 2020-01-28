package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"encoding/csv"
	"encoding/json"
)

type Author struct {
	Name string
	Date string
}

type Commit struct {
	Message string
	Author Author
}

type Branch struct {
	Name string
	Commit struct{
		Sha string
		Commit Commit
	}
}

func main() {
	branch := ParsePayload()
	fmt.Printf("%+v",branch)

	CreateAndWriteCsv(branch)
}

func ParsePayload() Branch{
	var branch Branch

	bodyBytes := callAPI()

	err := json.Unmarshal(bodyBytes, &branch)
	
	if err != nil {
		fmt.Println("error:", err)
	}
	
	return branch
}

func callAPI() []byte{
	accessToken := "6347414d9ed7eccfff96b5cff1d7b5458bdf0ff3"
	token := "token " + accessToken

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	query, _ := http.NewRequest("GET", "https://api.github.com/repos/diegogl12/github_web_hulk/branches/master", nil)

	parameters := query.URL.Query()

	query.Header.Add("Authorization", token)
	query.URL.RawQuery = parameters.Encode()

	result, _ := client.Do(query)

	bodyBytes, _ := ioutil.ReadAll(result.Body)

	return bodyBytes
}

// var data = [][]string{{"Repository", "Hello Readers of"}, {"TimeStamp", "golangcode.com"}}

func CreateAndWriteCsv(branch Branch) {
    file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

	branchString := toString(branch)

	err = writer.Write(branchString)
	checkError("Cannot write to file", err)
}

func toString(branch Branch) []string{
	var response = []string {branch.Name, branch.Commit.Sha, branch.Commit.Commit.Message, branch.Commit.Commit.Author.Name, branch.Commit.Commit.Author.Date}
	return response
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
