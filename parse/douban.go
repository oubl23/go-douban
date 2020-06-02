package parse

import(
	"log"
	"regexp"
	"strings"
	"strconv"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type DoubanMovie struct {
	Title string
	Subtitle string
	Other string
	Desc string
	Year string
	Area string
	Tag string
	Star string
	Comment string
	Quote string
}

type Page struct {
	Page int
	Url string
}

func GetPages(url string) []Page{
	client:= &http.Client{}
	req,_ := http.NewRequest("GET",url,nil)
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36 QIHU 360SE")
	res, _ := client.Do(req)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return ParsePages(doc)
}

func ParsePages(doc *goquery.Document) (pages []Page){
	pages = append(pages, Page{Page: 1, Url:""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")
		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})
	return pages
}

func ParseMovies(url string) (movies []DoubanMovie) {
	client:= &http.Client{}
	req,_ := http.NewRequest("GET",url,nil)
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36 QIHU 360SE")
	res, _ := client.Do(req)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()

		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		movie := DoubanMovie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		log.Printf("i: %d, movie: %v", i, movie)

		movies = append(movies, movie)
	})

	return movies
}
