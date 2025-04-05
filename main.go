package main

import (
	"fmt"
	"os"

	"github.com/filippixavier/gatorrss/internal/config"
)

func main() {
	conf, err := config.Read()

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	conf.SetUser("gueltir")

	newConf, err := config.Read()

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else {
		fmt.Println(newConf)
	}
}
