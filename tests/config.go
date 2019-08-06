package tests

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// --------------------------| config |-------------------------- \

// Config is a generaor configuraton
type Config struct {
	Host   string `yaml:"host"`   //!< storage's host
	Port   int    `yaml:"port"`   //!< storage's port
	Prefix string `yaml:"prefix"` //!< url prefix (e.g. /kv)
}

// ReadConfig creates a new config from the file
func ReadConfig(path string) Config {
	cfg := Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		err := fmt.Errorf("cannot read file:. %s", path)
		panic(err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		err = fmt.Errorf("cannot unmarshal file '%s':. %v", path, err)
		panic(err)
	}

	return cfg
}

// -----------|

// GetURL returns a ubs URL to a testing storage
func (cf *Config) GetURL(relative ...string) string {
	path := cf.Host + ":" + strconv.Itoa(cf.Port) + "/" + cf.Prefix + "/" + strings.Join(relative, "/")
	path = strings.ReplaceAll(path, "//", "/")
	return path
}
