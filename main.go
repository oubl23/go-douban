package main

import (
	"strings"
	"fmt"
	"github.com/oubl23/go-douban/parse"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

func Start(){
	var moives []parse.DoubanMovie
	pages := parse.GetPages(BaseUrl)

	for _,page := range pages{
		url := strings.Join([]string{BaseUrl, page.Url}, "");
		moives = append(moives, parse.ParseMovies(url)...)
	}
	fmt.Println(moives)
}

func main(){
	Start()
}
