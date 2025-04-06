package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/filippixavier/gatorrss/internal/database"
	"github.com/google/uuid"
)

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

func handlerReset(s *state, cmd command) error {
	if err := s.db.ClearUsers(context.Background()); err != nil {
		return err
	}

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if users, err := s.db.GetUsers(context.Background()); err != nil {
		return err
	} else {
		for _, user := range users {
			fmt.Printf("* %s", user.Name)
			if user.Name == s.cfg.CurrentUserName {
				fmt.Print(" (current)")
			}
			fmt.Println()
		}
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("missing parameters, expected: addfeed name url")
	}

	usr, e := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if e != nil {
		return e
	}

	if feed, e := s.db.CreateFeeds(context.Background(), database.CreateFeedsParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0], Url: cmd.args[1], UserID: usr.ID}); e != nil {
		return e
	} else {
		fmt.Println(feed)
	}

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("name: %s\n", feed.Name)
		fmt.Printf("url: %s\n", feed.Url)
		fmt.Printf("owner: %s\n", feed.Owner)
		fmt.Println()
	}

	return nil
}
