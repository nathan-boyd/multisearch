package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type SearchTemplate struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

type SearchTemplateContainer struct {
	Name            string           `json:"name"`
	SearchTemplates []SearchTemplate `json:"search_templates"`
}

type SearchTemplateCollection struct {
	Items []SearchTemplateContainer `json:"search_template_collections"`
}

func GetDefaultConfigFile() string {
	var err error
	var homeDir string
	if homeDir, err = os.UserHomeDir(); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.multisearch.json", homeDir)
}

func GetTemplateCollection(cfgFile string) (err error, collection SearchTemplateCollection) {
	if len(cfgFile) == 0 {
		cfgFile = GetDefaultConfigFile()
	}
	var jsonFile *os.File
	defer jsonFile.Close()
	if jsonFile, err = os.Open(cfgFile); err != nil {
		return
	}
	var byteValue []byte
	if byteValue, err = ioutil.ReadAll(jsonFile); nil != err {
		return
	}
	if err = json.Unmarshal(byteValue, &collection); nil != err {
		return
	}
	return
}
