package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"strings"
	"fmt"
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

	csvFile, err := os.Open("./beerstyles.csv")
	if err != nil {
		log.Fatal("File: ", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, csverr := reader.ReadAll()
	if csverr != nil {
		log.Fatal("CSV: ", err)
	}

	styles := []Style{}

	for _, row := range records {
		log.Println(row[0], row[4])
		styles = append(styles, Style{
			Name:    row[0],
			Desc:    strings.ReplaceAll(row[2], "<br>", "\n\n"),
			Details: row[3],
			Slug:    slug.Make(row[4]),
		})
	}

	mdfile, mderr := os.ReadFile("./style.mdx")
	if mderr != nil {
		log.Fatal("MDX: ", mderr)
	}

	tmpl, tmplerr := template.New("mdx").Parse(string(mdfile))
	if tmplerr != nil {
		log.Fatal(tmplerr)
	}

	for i, style := range styles[1:] {

		stylefilename := fmt.Sprintf("s_%d.mdx", i)
		stylefile, styleerr := os.Create("./src/content/style-guide/" + stylefilename)
		if styleerr != nil {
			log.Fatal("Write: ", styleerr)
			break
		}

		oserr := tmpl.Execute(stylefile, style)
		if oserr != nil {
			log.Fatal(oserr)
		}

		/*
		if i == 1 {
			break
		}
		*/
	}
}
