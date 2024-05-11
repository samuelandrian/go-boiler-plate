package config

import (
	"flag"
	stringhelper "go-boiler-plate/pkg/stringhelper"
	"log"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v2"
)

func (config *AppConfig) SetEnvironment() {
	env := flag.String("env", "", "env Mode project local, staging, production")
	flag.Parse()
	config.App.Env = stringhelper.StringPointerToString(env)
}

func (config *AppConfig) ReadConfigFromFile() {
	var filepath string
	_, filename, _, _ := runtime.Caller(1)
	switch TypeEnvironment(config.App.Env) {
	case LOCAL:
		filepath = path.Join(path.Dir(path.Dir(filename)), "/pkg/config/local.yml")
	default:
		filepath = path.Join(path.Dir(path.Dir(filename)), "/pkg/config/config.yml")
	}
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

}
