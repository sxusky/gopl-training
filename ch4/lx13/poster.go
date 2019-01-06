package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const omdbURL = "https://omdbapi.com/?apikey=BanMePlz"

type movieSearchResult struct {
	Response string
	Poster   string
	Error    string
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Printf("useage: %s <year> <moive titel>\n", os.Args[0])
		os.Exit(0)
	}
	year := os.Args[1]
	title := os.Args[2]
	downloadMoviePoster(year, title)
}

func downloadMoviePoster(year string, title string) {
	result, err := obtainMovieInfomation(year, title)
	if err != nil {
		log.Fatal(err)
	}
	if result.Response == "True" {
		filename := year + "." + strings.Replace(title, " ", "_", -1) + ".jpg"
		if err := download(result.Poster, filename); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(result.Error)
	}
}

func obtainMovieInfomation(year string, title string) (*movieSearchResult, error) {
	q := url.QueryEscape(title)
	url := omdbURL + "&y=" + year + "&t=" + q
	fmt.Printf(url + "\n")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result movieSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func download(fromURL string, toFile string) error {
	fmt.Printf(fromURL + "\n")
	resp, err := http.Get(fromURL)
	if err != nil {
		return err
	}
	file, err := os.Create(toFile)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
