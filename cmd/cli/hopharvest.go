package main

/*
TODO: split %
*/
import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/google/uuid"
)

type DataSet struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BrewData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Hop struct {
	Name         string     `json:"name"`
	Data         []DataSet  `json:"data,omitempty"`
	BrewData     []BrewData `json:"brewdata,omitempty"`
	Substitution []Hop      `json:"substitutes,omitempty"`
	Uuid         string     `json:"uuid"`
}

var Hops []Hop

func HopHarvest() {
	log.Println("Harvesting hops...")

	// Colly
	c := colly.NewCollector(
		colly.CacheDir("./cache"),
		colly.MaxDepth(1),
		colly.AllowedDomains("beermaverick.com"),
	)
	c.Limit(&colly.LimitRule{
		Delay: 4 * time.Second,
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		l := e.Attr("href")

		if strings.Contains(l, "/hop/") {
			c.Visit(e.Request.AbsoluteURL(l))
		}
	})

	c.OnHTML(".single-entry", func(e *colly.HTMLElement) {
		// Only if brewvalues is present
		if len(e.DOM.Find(".brewvalues").Nodes) > 0 {
			// New hop
			hop := Hop{Uuid: uuid.New().String()}
			// First element is header
			hop.Name = e.DOM.Find(".entry-title").First().Text()

			// Table
			e.DOM.Find(".wp-block-table tr").Each(func(_ int, s *goquery.Selection) {
				f := s.Find("th").Text()
				l := s.Find("td").Text()

				hop.Data = append(hop.Data, DataSet{
					Name:  strings.TrimSpace(f),
					Value: strings.TrimSpace(l),
				})
			})
			// Entries
			e.DOM.Find("h2").Each(func(i int, s *goquery.Selection) {
				//
				h := s.Text()
				p := ""
				// Hop substitutes p/ul
				if strings.Contains(h, "Substitutions") {
					s.NextAllFiltered("p").Each(func(_ int, t *goquery.Selection) {
						p = p + t.Text() + " "
					})
					s.NextAllFiltered("ul").Find("li").Each(func(_ int, l *goquery.Selection) {
						hop.Substitution = append(hop.Substitution, Hop{
							Name: strings.TrimSpace(l.Text()),
						})
					})
				} else {
					s.NextAllFiltered("p").Each(func(_ int, t *goquery.Selection) {
						p = p + t.Text() + " "
					})
				}

				hop.Data = append(hop.Data, DataSet{
					Name:  strings.TrimSpace(h),
					Value: strings.TrimSpace(p),
				})
			})
			// Brew values
			e.DOM.Find(".brewvalues tr").Each(func(_ int, s *goquery.Selection) {
				hop.BrewData = append(hop.BrewData, BrewData{
					Name:  s.Find("th").Text(),
					Value: s.Find("td").Text(),
				})
			})

			Hops = append(Hops, hop)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		// log.Println("Visiting ", r.URL.String())
	})

	c.Visit("https://beermaverick.com/hops/")

	c.Wait()

	// Write to file
	b, _ := json.MarshalIndent(Hops, "", "\t")

	f, err := os.Create("./hops.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, wr := f.Write(b); wr != nil {
		panic(wr)
	}
}
