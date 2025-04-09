-- name: CreateFeeds :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, f.url, u.name owner
FROM feeds f
INNER JOIN users u ON f.user_id = u.id;

-- name: GetFeedByUrl :one
SELECT *
FROM feeds
WHERE url = $1
LIMIT 1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = $2, updated_at = $2
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at
NULLS FIRST;