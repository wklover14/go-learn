package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Query struct {
		Pages map[string]struct {
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}

func main() {
	word := "chat" // Remplacez ceci par le mot que vous recherchez
	url := fmt.Sprintf("https://fr.wiktionary.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=%s", word)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête GET:", err)
		return
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Println("Erreur lors de l'analyse de la réponse JSON:", err)
		return
	}

	for _, page := range apiResponse.Query.Pages {
		fmt.Println(page.Extract)
	}
}
