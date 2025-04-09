package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/filippixavier/gatorrss/internal/config"
	"github.com/filippixavier/gatorrss/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print("%w\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbURL)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	st := state{cfg: &cfg, db: dbQueries}

	cmds := commands{list: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogging)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))

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
