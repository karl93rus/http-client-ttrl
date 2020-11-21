package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/karl93rus/http-client/client"
)

var (
	gc = getGitClient()
)

func getGitClient() client.HttpClient {
	gc := client.NewClient()
	headers := make(http.Header)
	headers.Set("Accept", "application/json")

	gc.SetHeaders(headers)

	return gc
}

func getGit() {
	gitHeaders := make(http.Header)
	gitHeaders.Set("Accept", "application/xml")
	gitHeaders.Set("Authorization", "Bearer ABC-123")

	res, err := gc.Get("https://api.github.com", gitHeaders)
	if err != nil {
		panic(err)
	}

	bytes := res.Body
	r, _ := ioutil.ReadAll(bytes)

	fmt.Println(string(r))

}

func main() {
	getGit()
	getGit()
	getGit()
}
