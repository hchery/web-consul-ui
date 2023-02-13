package app

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

/**********************************
 * Date: 2023/2/3
 * Author: hchery
 * Home: https://github.com/hchery
 *********************************/

func fOpen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot open configuration file: %s, cause: %s", path, err.Error())
		os.Exit(ExitWithConfigurationFileError)
	}
	return f
}

func fClose(path string, f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Printf("Cannot release configuration file: %s, cause: %s", path, err.Error())
	}
}

func unmarshal(path string, f *os.File, v any) {
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(v); err != nil {
		fmt.Printf("Cannot parse configuration file: %s, cause: %s", path, err.Error())
		os.Exit(ExitWithConfigurationFileError)
	}
}

func YamlUnmarshal(path string, v any) {
	yml := fOpen(path)
	defer fClose(path, yml)
	unmarshal(path, yml, v)
}

const (
	SnowflakeConf = "conf/snowflake.yaml"
)
