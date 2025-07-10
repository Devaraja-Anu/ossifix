package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

const (
	DirPerm  = 0755
	FilePerm = 0644
)

func main() {

	dirStructure := []string{
		"testapp/cmd/api",
		"testapp/internal",
		"testapp/migrations",
	}

	for _, dir := range dirStructure {
		if err := os.MkdirAll(dir, DirPerm); err != nil {
			panic(err)
		}
	}

	tmpl, err := template.ParseFiles("templates/rest/main.go.tmpl")
	if err != nil {
		fmt.Println("Error parsing templ")
		log.Fatal(err)
	}

	file, err := os.Create("testapp/cmd/api/main.go")
	if err != nil {
		fmt.Println("Error creating file templ")
		log.Fatal(err)
	}

	err = tmpl.Execute(file, nil)

	if err != nil {
		fmt.Println("Error executing templ")
	}

	fmt.Println("Success")

}
