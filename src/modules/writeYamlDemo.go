package modules

//
//import (
//	"fmt"
//	"gopkg.in/yaml.v2"
//	"log"
//	"os"
//)
//
//// YamlConfig is exported.
//type yamlConfig struct {
//	Connection struct {
//		Hits int64 `yaml:"hits"`
//		Time int64 `yaml:"time"`
//	}
//}
//
//func WriteYml() bool {
//
//	documentToSave := yamlConfig{}
//	documentToSave.Connection.Time = 41
//	documentToSave.Connection.Hits = 42
//
//	d, err := yaml.Marshal(&documentToSave)
//	if err != nil {
//		log.Fatalf("Failed to Save that sexy yaml")
//	}
//
//	fmt.Println(string(d))
//
//	f, err := os.OpenFile("testWrite.yaml", os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer f.Close()
//
//	_, err = f.Write(d)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return true
//}
//
//
