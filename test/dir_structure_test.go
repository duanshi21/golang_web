package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type DirConfig struct {
	Name     string      `json:"name"`
	Children []DirConfig `json:"children,omitempty"`
}

func createDirectoryStructure(config DirConfig, parentPath string) error {
	if config.Name == "" {
		fmt.Printf("Empty directory name: %+v\n", config)
		parentPath = ".."
	}

	dirPath := fmt.Sprintf("%s/%s", parentPath, config.Name)
	fmt.Println(dirPath)
	if err := os.Mkdir(dirPath, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	for _, child := range config.Children {
		if err := createDirectoryStructure(child, dirPath); err != nil {
			return err
		}
	}

	return nil
}

func TestGenerate(t *testing.T) {
	jsonFile, err := os.Open("dir_structure.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	var rootDir DirConfig
	if err := json.NewDecoder(jsonFile).Decode(&rootDir); err != nil {
		panic(err)
	}

	if err := createDirectoryStructure(rootDir, "/"); err != nil {
		panic(err)
	}
}
