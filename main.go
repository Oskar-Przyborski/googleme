package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/savioxavier/termlink"
)

type ResponseType struct {
	Items []ResponseItem `json:"items"`
}
type ResponseItem struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	DisplayLink string `json:"displayLink"`
	Snippet     string `json:"snippet"`
}

const API_KEY string = "AIzaSyAf4AXGBTAJNkk4i2_iUctk0sU-plFrNdA"
const BROWSER_ID string = "24ff328fa6314445a"

func getSearchUrl(query string) string {
	return fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s", API_KEY, BROWSER_ID, url.QueryEscape(query))
}

func main() {
	query := strings.Join(os.Args[1:], " ")
	url := getSearchUrl(query)
	println("Searching for:", "\033[36m"+query+"\033[0m")
	println()
	response, err := http.Get(url)
	if err != nil {
		println("Error while getting response from google:\n" + err.Error() + "\nMaybe you don't have internet acces or the app's api keys don't work")
		return
	}

	var data ResponseType
	json.NewDecoder(response.Body).Decode(&data)

	if len(data.Items) == 0 {
		println("Found 0 results ://")
	} else {
		for _, item := range data.Items {
			println(item.DisplayLink)
			println(termlink.ColorLink(item.Title, item.Link, "green"))
			println(item.Snippet)
			println()
			println()
		}
	}
}
