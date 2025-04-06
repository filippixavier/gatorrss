package main

import (
	"fmt"

	"github.com/filippixavier/gatorrss/internal/config"
	"github.com/filippixavier/gatorrss/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if clb, exists := c.list[cmd.name]; exists {
		return clb(s, cmd)
	} else {
		return fmt.Errorf("command doesn't exist")
	}
}
