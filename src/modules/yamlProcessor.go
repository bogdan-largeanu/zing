package modules

import (
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

func (yamlStruc *YamlStructure) ReadYml() *YamlStructure {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, yamlStruc)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	checkYamlFormat(yamlStruc)

	return yamlStruc
}

func checkYamlFormat(yamlStruct *YamlStructure) bool {
	//todo find why this is read twice
	//todo check key and bash file are valid throw stout error otherwise to use
	for i := range *yamlStruct {
		println("DEBUG", yamlStruct[i])
	}
	return false
}
