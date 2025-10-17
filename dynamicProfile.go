package main

import (
	"net/http"
	"time"
	"encoding/json"
	"io"
	"fmt"
	"log"
)


const url = "https://catfact.ninja/fact"

type User struct {
	Email  string 
	Name   string
	Stack  string
}

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type Profile struct {
	Status     string     `json:"status"`
	User       User       `json:"user"`
	Timestamp  time.Time  `json:"timestamp"`
	Fact       string     `json:"fact"`
}



func factHandler(w http.ResponseWriter, req *http.Request) {
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Failed to read response: %v", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var fact CatFact
	if err := json.Unmarshal([]byte(body), &fact); err != nil {
		return 
	}

	profile := Profile {
		Status: "success",
		User: User {
			Email: "mwangiruoya@gmail.com",
			Name: "Victor Mwangi",
			Stack: "Go/Gin",
		},
		Timestamp: time.Now().UTC(),
		Fact: fact.Fact,
	}

	data, _ := json.MarshalIndent(profile, "", " ")
	w.Header().Set("Content-Type", "application/json")

	w.Write(data)

}


func main() {
	http.HandleFunc("/me", factHandler)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
