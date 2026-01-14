package main

// use for generate duplicate module with to be complied
// with same architecture

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// use cli to generate template
// ex. go run cmd/genmodule/main.go product
func main() {
	name := os.Args[1]

	data := map[string]string{
		"Name":   cases.Title(language.Und).String(name),
		"Module": name,
	}

	templateDir, err := os.ReadDir("internal/template")
	if err != nil {
		panic(err)
	}

	for _, file := range templateDir {
		templatePath := fmt.Sprintf("internal/template/%s", file.Name())
		tpl, err := template.ParseFiles(templatePath)
		if err != nil {
			panic(err)
		}
		err = os.MkdirAll("internal/domain/"+name, 0755)
		if err != nil {
			panic(err)
		}
		outputPath := fmt.Sprintf("internal/domain/%s/%s", name, file.Name())
		if before, ok := strings.CutSuffix(outputPath, ".tpl"); ok {
			outputPath = before
		}
		f, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := tpl.Execute(f, data); err != nil {
			panic(err)
		}
	}

}
