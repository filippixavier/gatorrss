package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/filippixavier/gatorrss/internal/config"
	"github.com/filippixavier/gatorrss/internal/database"
	"github.com/google/uuid"
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

func handlerLoging(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("missing the username argument")
	}

	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		return errors.New(cmd.args[0] + "  doesn't exist, please register it first!")
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("Username has been set to %s\n", cmd.args[0])

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("missing the username argument")
	}

	if u, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		return errors.New(u.Name + " already exist")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0]})
	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)

	fmt.Printf("user %s has been successfully created\n", cmd.args[0])

	fmt.Printf("%v", user)

	return nil
}
