package scaffold

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

const (
	DirPerm  = 0755
	FilePerm = 0644
)

func CreateFiles(rootpath string) error {

	dirStructure := []string{
		filepath.Join(rootpath, "cmd", "api"),
		filepath.Join(rootpath, "internal"),
		filepath.Join(rootpath, "migrations"),
	}

	for _, dir := range dirStructure {
		if err := os.MkdirAll(dir, DirPerm); err != nil {
			return err
		}
	}

	tmpl, err := template.ParseFiles(filepath.Join("templates", "rest", "main.go.tmpl"))
	if err != nil {
		fmt.Println("Error creating file templ")
		return err
	}

	file, err := os.Create(filepath.Join(rootpath, "cmd", "api", "main.go"))
	if err != nil {
		fmt.Println("Error creating file templ")
		return err
	}

	err = tmpl.Execute(file, nil)

	if err != nil {
		fmt.Println("Error executing templ")
		return err
	}

	return nil
}
