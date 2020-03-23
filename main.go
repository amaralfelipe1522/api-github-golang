package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Usuario armazena a estrutura base de dados do usuario
type Usuario struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	QtdRepos int    `json:"public_repos"`
	Bio      string `json:"bio"`
	URL      string `json:"html_url"`
}

func get(url string) {
	fmt.Println("Starting the application...")
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario Usuario
	json.Unmarshal(responseData, &usuario)

	fmt.Println(usuario)
	defer fmt.Println("Terminating the application...")
}

func main() {
	//baseURL := "https://api.github.com"
	userURL := "https://api.github.com/users/amaralfelipe1522"
	get(userURL)
}
