package services

import (
	"fmt"
	"encoding/json"
	"../structs"
	"net/http"
	"io/ioutil"
)

func ParsePayload(owner string, repo string) structs.Branch{
	var branch structs.Branch

	bodyBytes := callAPI(owner, repo)

	err := json.Unmarshal(bodyBytes, &branch)
	
	if err != nil {
		fmt.Println("error:", err)
	}
	
	return branch
}

func callAPI(owner string, repo string) []byte{
	accessToken := "e5688e0f8051160379fb860856878963fb7dd254"
	token := "token " + accessToken

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/branches/master"
	query, _ := http.NewRequest("GET", url, nil)

	parameters := query.URL.Query()

	query.Header.Add("Authorization", token)
	query.URL.RawQuery = parameters.Encode()

	result, _ := client.Do(query)

	bodyBytes, _ := ioutil.ReadAll(result.Body)

	return bodyBytes
}
