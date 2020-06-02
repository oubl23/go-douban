package main

import (
	"fmt"
	"github.com/oubl23/go-douban/parse"
	"strings"
	"time"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

func Start() {
	var movies []parse.DoubanMovie
	pages := parse.GetPages(BaseUrl)
	ch := make(chan []parse.DoubanMovie)
	for _, page := range pages {
		url := strings.Join([]string{BaseUrl, page.Url}, "")
		go func(ch chan []parse.DoubanMovie) {
			ch <- parse.ParseMovies(url)
		}(ch)
	}

	for range pages {
		movies = append(movies, <-ch...)
	}
	close(ch)
}

func main() {
	startTime := time.Now().UnixNano()
	Start()
	endTime := time.Now().UnixNano()
	fmt.Println("time spend", float64((endTime-startTime)/1e6), "ms")
}
