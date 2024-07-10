package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/gosimple/slug"
)

type Style struct {
	Name    string
	Desc    string
	Details string
	Slug    string
}

func CreateMdFiles() {
	log.Println("Generate MD files")

	csvFile, err := os.Open("./hopsstyle.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, csverr := reader.ReadAll()
	if csverr != nil {
		log.Fatal(err)
	}

	styles := []Style{}

	for _, row := range records {
		// log.Println(row[0], row[3], row[2])
		styles = append(styles, Style{
			Name:    row[0],
			Desc:    strings.ReplaceAll(row[3], "|", "\n\n"),
			Details: strings.ReplaceAll(row[2], "|", "\n\n"),
			Slug:    slug.Make(row[0]),
		})
	}

	mdfile, mderr := os.ReadFile("./style.mdx")
	if mderr != nil {
		log.Fatal(mderr)
	}

	tmpl, tmplerr := template.New("mdx").Parse(string(mdfile))
	if tmplerr != nil {
		log.Fatal(tmplerr)
	}

	for i, style := range styles {

		stylefilename := style.Slug + ".mdx"
		stylefile, styleerr := os.Create("./src/content/style-guide/" + stylefilename)
		if styleerr != nil {
			log.Fatal(styleerr)
			break
		}

		oserr := tmpl.Execute(stylefile, style)
		if oserr != nil {
			log.Fatal(oserr)
		}

		if i == 10 {
			break
		}
	}
}
