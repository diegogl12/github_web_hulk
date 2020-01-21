package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	accessToken := "897f60893614063ad9617d5f19d980f19a869a47"
	// accessToken := "57832d165bc579b697fe95befeadabb7ee66a004"
	token := "token " + accessToken

	// result, _ := http.Get("https://api.github.com/users/diegogl12/repos")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	query, _ := http.NewRequest("GET", "https://api.github.com/users/sumup/repos", nil)

	query.Header.Add("Authorization", token)
	result, _ := client.Do(query)

	bodyBytes, _ := ioutil.ReadAll(result.Body)
	bodyString := string(bodyBytes)

	fmt.Println(bodyString)
}
