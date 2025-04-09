package main

import (
	"context"
	"errors"

	"github.com/filippixavier/gatorrss/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

		if err != nil {
			return errors.New(s.cfg.CurrentUserName + "  doesn't exist, please register it first!")
		}

		return handler(s, c, usr)
	}
}
