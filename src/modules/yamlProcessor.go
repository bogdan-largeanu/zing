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
		LiteralBlockBashFile string `yaml:"literal_block_bash_file"`
	} `yaml:"run"`
}

func ReadYml() (YamlStructure, int, error) {
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

func checkYamlFormat(structure YamlStructure) (YamlStructure, int, error) {
	for i := range structure {
		if structure[i].Run.Key == "" {
			description := fmt.Sprintf("\n Description: %#v", structure[i].Run.Description)
			bashScript := fmt.Sprintf("\n literal_block_bash_file: %#v", structure[i].Run.LiteralBlockBashFile)
			path := fmt.Sprintf("\n path: %#v", structure[i].Run.Path)
			return structure, i, errors.New(description + bashScript + path +
				"\n ERROR -> Missing field \"run\" ")
		}

		if structure[i].Run.LiteralBlockBashFile == "" {
			description := fmt.Sprintf("\n Description: %#v", structure[i].Run.Description)
			run := fmt.Sprintf("\n run: %#v", structure[i].Run.Key)
			path := fmt.Sprintf("\n path: %#v", structure[i].Run.Path)
			return structure, i, errors.New(description + run + path +
				"\n ERROR -> Missing field \"literal_block_bash_file\" ")
		}

	}
	return structure, 0, nil
}
