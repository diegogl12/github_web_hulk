package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"encoding/csv"
	"encoding/json"
	"regexp"
)

func main() {
	accessToken := "dc1a06bf8a54ff50ecbd1c38e0208e1b86093d28"
	token := "token " + accessToken

	// result, _ := http.Get("https://api.github.com/users/diegogl12/repos")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	query, _ := http.NewRequest("GET", "https://api.github.com/users/diegogl12/repos", nil)

	parameters := query.URL.Query()
	parameters.Add("sha","master")

	query.Header.Add("Authorization", token)
	query.URL.RawQuery = parameters.Encode()

	result, _ := client.Do(query)

	bodyBytes, _ := ioutil.ReadAll(result.Body)
	// bodyString := string(bodyBytes)

	type Commit struct {
		Sha string
		Node_id string
	}

	var commits []Commit

	err := json.Unmarshal(bodyBytes, &commits)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v",commits)

	// CreateAndWriteCsv()

	// var validURL = regexp.MustCompile(`repos\/(.*?)\/`)
	// fmt.Println(validURL.Find([]byte(`https://api.github.com/repos/diegogl12/github_web_hulk/commits/550f0f858ae776b94d00b7729b1d0d79164e9c8e`)))
}


var data = [][]string{{"Repository", "Hello Readers of"}, {"TimeStamp", "golangcode.com"}}

func CreateAndWriteCsv() {
    file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range data {
        err := writer.Write(value)
        checkError("Cannot write to file", err)
    }
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
