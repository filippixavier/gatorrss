# Gatorrss

Small RSS aggregator from a boot.dev guided project,
A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
go install "github.com/filippixavier/gatorrss@latest"
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
gatorrss register <name>
```

Add a feed:

```bash
gatorrss addfeed <url>
```

Start the aggregator:

```bash
gatorrss agg 30s
```

View the posts:

```bash
gatorrss browse [limit]
```

There are a few other commands you'll need as well:

- `gatorrss login <name>` - Log in as a user that already exists
- `gatorrss users` - List all users
- `gatorrss feeds` - List all feeds
- `gatorrss follow <url>` - Follow a feed that already exists in the database
- `gatorrss unfollow <url>` - Unfollow a feed that already exists in the database
