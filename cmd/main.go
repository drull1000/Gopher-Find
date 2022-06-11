package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sherlockgo/cmd/color"
	"strings"
)

func main() {
	file, err := os.Open("./resources/data.json")
	handleError(err)
	var endpoints map[string]interface{}
	err = json.NewDecoder(file).Decode(&endpoints)
	handleError(err)
	fmt.Print("Type the username: ")
	username := getInput()

	for websiteName, parameter := range endpoints {
		websiteURL := parameter.(map[string]interface{})["url"]
		checkURL(websiteURL, websiteName, username)
	}
	defer file.Close()
}

func checkURL(websiteURL interface{}, websiteName interface{}, username string) {
	url := formatedURL(websiteURL.(string), username)
	resp, err := http.Get(url)
	handleConnectionError(err)
	defer resp.Body.Close()
	handleError(err)
	if resp.StatusCode == 200 {
		fmt.Println(color.Green+"[+] FOUND -", websiteName, color.Reset)
		fmt.Println(url)
	} else {
		fmt.Println(color.Red+"[-] NOT FOUND -", websiteName, color.Reset)
	}
}
func formatedURL(url string, username string) string {
	return strings.Replace(url, "{}", username, -1)
}
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
func handleConnectionError(err error) {
	//TODO: make code continue to run after error message is shown.
	if err != nil {
		fmt.Println(color.Red+"[-] CONNECTION ERROR -", err, color.Reset)
	}
}

func getInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}
