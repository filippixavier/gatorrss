package main

import (
	"fmt"
	"os"

	"github.com/filippixavier/gatorrss/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print("%w\n", err)
		os.Exit(1)
	}

	st := state{cfg: &cfg}

	cmds := commands{list: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLoging)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Missing arguments")
		os.Exit(1)
	}

	cmd := command{name: args[1], args: args[2:]}

	if err := cmds.run(&st, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
