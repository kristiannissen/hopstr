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
		})
	}

	mdfile, mderr := os.ReadFile("./style.md")
	if mderr != nil {
		log.Fatal(mderr)
	}

	for i, style := range styles {
		tmpl, tmplerr := template.New("md").Parse(string(mdfile))
		if tmplerr != nil {
			log.Fatal(tmplerr)
			break
		}

		stylefilename := slug.Make(style.Name) + ".md"
		stylefile, styleerr := os.Create("./src/pages/style-guide/" + stylefilename)
		if styleerr != nil {
			log.Fatal(styleerr)
			break
		}

		oserr := tmpl.Execute(stylefile, style)
		if oserr != nil {
			log.Fatal(oserr)
		}

		if i == 2 {
			break
		}
	}
}
