package main

import (
	"flag"
	"fmt"

	"github.com/TheMrViper/gini"
)

var (
	configName   = flag.String("configName", "config.ini", "Name of you config")
	configCreate = flag.Bool("configCreate", false, "Create empty config from structure")
)

// used as config for package1
type SubConfig1 struct {
	Field int `ini-default:"-1234"` // set default value if this field not in config.ini
}

// Used as main app config
type Config struct {

	// add main config fields
	MainField string `ini-name:"String"` // add other name
	IgnoreMe  string `ini:"-"`           // ignore this field

	// include other package configs
	PkgConfig1 SubConfig1 `ini-name:"PkgConfig"` // read data from {PkgConfig} section
	PkgConfig2 SubConfig1 `ini-name:"PkgConfig"` // read data from {PkgConfig} section

	// use same structure, but different values
	PkgConfig3 SubConfig1 `ini-name:"PkgConfig1"` // read data from {PkgConfig1} section

	SubPkgConfig1 SubConfig2 // add other package
	SubPkgConfig2 SubConfig3 // add other package
}

// used as config for package2
type SubConfig2 struct {
	Field uint `ini-default:"1234"` // set default value if this field not in config.ini

	PkgConfig SubConfig1 `ini-name:"PkgConfig"` // read data from {PkgConfig} section
}

// used as config for package3
type SubConfig3 struct {
	Field bool `ini-default:"true"` // set default value if this field not in config.ini

	PkgConfig SubConfig1 `ini-name:"PkgConfig1"` // read data from {PkgConfig1} section
}

func main() {
	flag.Parse()

	config := Config{}
	if *configCreate {
		if err := gini.WriteConfig(*configName, &config); err != nil {
			fmt.Println("Error: ", err)
		}
		return
	}

	if err := gini.ReadConfig("config.ini", &config); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Config: ", config)
}
