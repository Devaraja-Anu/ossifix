package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	DirPerm  = 0755
	FilePerm = 0644
)

func CreateScaffold(rootpath string, router string) error {

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

	tmpls, err := newTemplateCache("/rest/common")
	if err != nil {
		fmt.Println("Error executing templ")
	}
	err = createAndWriteFiles(rootpath, tmpls)
	if err != nil {
		fmt.Println("error writing files")
	}

	tmpls, err = newTemplateCache(fmt.Sprintf("/rest/routers/%s", strings.ToLower(router)))
	if err != nil {
		fmt.Println("Error executing templ")
	}
	err = createAndWriteFiles(rootpath, tmpls)
	if err != nil {
		fmt.Println("error writing files")
	}

	return nil
}
