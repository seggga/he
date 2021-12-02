package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Timeout int `yaml:"timeout"`
	// MaxResults   int    `yaml:"maxresults"`
	// MaxErrors    int    `yaml:"maxerrors"`
	// Url          string `yaml:"url"`
	// ReqTimeout   int    `yaml:"reqtimeout"`
	// CrawlTimeout int    `yaml:"crawltimeout"`
}

func main() {
	// get application config
	cfg, err := ReadConfig()
	if err != nil {
		fmt.Printf("Unable to read config.yaml, %v.\nProgram exit", err)
		// log.Errorf()
	}
	fmt.Printf("%+v\n", cfg)

	// obtain users query
	fmt.Print("Enter the query: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading query, %v", err)
		return
	}
	// convert CRLF to LF
	query = strings.Replace(query, "\n", "", -1)
	// log.Debug()
	fmt.Println(query)

}

// ReadConfig returns a structure with data from config-file
func ReadConfig() (*Config, error) {
	// read config file
	configData, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		return nil, err
	}
	// decode config
	cfg := new(Config)
	err = yaml.Unmarshal(configData, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
