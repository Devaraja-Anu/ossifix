package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// enter relevant folders from template eg:/rest/common ,etc
func newTemplateCache(folder string) (map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)
	files, err := filepath.Glob(fmt.Sprintf("./templates%s/*.tmpl", folder))
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		fileName := filepath.Base(file)
		ts, err := template.New(fileName).ParseFiles(file)
		if err != nil {
			return nil, err
		}
		cache[fileName] = ts
	}

	return cache, nil
}

func createAndWriteFiles(rootpath string, tmpls map[string]*template.Template) error {
	for _, tmpl := range tmpls {
		fileName := strings.TrimSuffix(tmpl.Name(), ".tmpl")
		file, err := os.Create(filepath.Join(rootpath, "cmd", "api", fileName))
		if err != nil {
			fmt.Println("Error creating file templ")
			return err
		}
		err = tmpl.Execute(file, nil)

		file.Close()
	}
	return nil
}
