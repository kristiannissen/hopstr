package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func HarvestBeerStyles() {
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
			content = content + strings.TrimSpace(s.Text()) + "\n"
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

	c.Visit("")
}
