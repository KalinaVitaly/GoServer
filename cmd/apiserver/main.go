package main

import (
	"Diplom/internal/app/apifilesystem"
	"Diplom/internal/app/apiserver"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	configFileSystem := apifilesystem.NewConfigDirecories()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.DecodeFile(configPath, configFileSystem)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Config read : ", config, configFileSystem)

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
