package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(url string) {
	fmt.Println("Starting the application...")
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("Terminating the application...")
}

func main() {
	//baseURL := "https://api.github.com"
	userURL := "https://api.github.com/users/amaralfelipe1522/repos"
	get(userURL)
}
