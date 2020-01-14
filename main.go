package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	result, _ := http.Get("https://api.github.com")
	
	bodyBytes, _ := ioutil.ReadAll(result.Body)
  bodyString := string(bodyBytes)
		
	fmt.Println(bodyString)
}

