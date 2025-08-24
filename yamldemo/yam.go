package yamldemo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Dependency describes a dependency
type Dependency struct {
	Name          string
	Version       string
	RepositoryURL string `yaml:"repository"`
}

type YAMLFile struct {
	Dependencies []Dependency `yaml:"dependencies"`
}

func Run() {
	f, err := os.Open(filepath.Join("yamldemo", "data.yml"))
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)

	var yamlFile YAMLFile
	err = dec.Decode(&yamlFile)
	if err != nil {
		log.Fatalf("dec.Decode() failed with '%s'\n", err)
	}

	fmt.Printf("Decoded YAML dependencies: %#v\n", yamlFile.Dependencies)
}
