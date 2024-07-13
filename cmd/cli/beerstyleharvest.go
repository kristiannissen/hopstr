package main

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func BeerStyleHarvest() {
	log.Println("Harvesting...")
	//
	f, err := os.Create("./styleguide.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	c := colly.NewCollector()

	c.OnHTML(".beer-style", func(e *colly.HTMLElement) {
		heading := e.DOM.Find(".list-with-heading>li:first-child").Text()

		list := e.DOM.Find(".list-with-heading>li:not(:first-child)")

		var content string
		list.Each(func(_ int, s *goquery.Selection) {
			// content = content + strings.TrimSpace(s.Text()) + "\n"
			//
			pairs := strings.Split(s.Text(), ":")

			re := regexp.MustCompile(`^[[:space:]]|`)
			if re.Match([]byte(pairs[1])) {
				pairs[1] = string(re.ReplaceAll([]byte(pairs[1]), []byte("")))
			}

			content = content + strings.TrimSpace(pairs[0]) + ": " + strings.TrimSpace(pairs[1]) + "|"
		})

		var brew string
		details := e.DOM.Find(".horizontal.wider")
		details.Find("li").Each(func(_ int, s *goquery.Selection) {
			brew = brew + s.Text() + "\n"
		})

		writer.Write([]string{heading, content, brew})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.brewersassociation.org/edu/brewers-association-beer-style-guidelines/")
}
