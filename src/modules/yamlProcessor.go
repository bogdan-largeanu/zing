package modules

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type YamlStructure []struct {
	Run struct {
		Key                  string `yaml:"key"`
		Description          string `yaml:"description"`
		Path                 string `yaml:"path"`
		LiteralBlockBashFile string `yaml:"bash_scripts"`
	} `yaml:"run"`
}

func ReadYml() (YamlStructure, error) {
	var yamlStruc YamlStructure

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &yamlStruc)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return checkYamlFormat(yamlStruc)
}

func checkYamlFormat(structure YamlStructure) (YamlStructure, error) {
	for i := range structure {
		if structure[i].Run.Key == "" {
			description := fmt.Sprintf("\n Description: %#v", structure[i].Run.Description)
			bashScript := fmt.Sprintf("\n bash_scripts: %#v", structure[i].Run.LiteralBlockBashFile)
			path := fmt.Sprintf("\n path: %#v", structure[i].Run.Path)
			return structure, errors.New(description + bashScript + path +
				"\n ERROR -> Missing field \"run\" ")
		}

		if structure[i].Run.LiteralBlockBashFile == "" {
			description := fmt.Sprintf("\n Description: %#v", structure[i].Run.Description)
			run := fmt.Sprintf("\n run: %#v", structure[i].Run.Key)
			path := fmt.Sprintf("\n path: %#v", structure[i].Run.Path)
			return structure, errors.New(description + run + path +
				"\n ERROR -> Missing field \"bash_scripts\" ")
		}

	}
	return structure, nil
}
