// config.go
package main

import (
	"flag"
	"fmt"
	"os"
)

var config string

func init() {
	flag.StringVar(&config, "c", "./config.json", "set http server config file path")
}

func initConfig() {
	if f, err := os.Open(config); err != nil {
		fmt.Println(err)
		// return err
	} else {
		defer f.Close()
		// s
	}
}
