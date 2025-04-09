-- name: CreateFeedsFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feeds_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_feed_follow.*,
    feeds.name feed_name,
    users.name user_name
FROM inserted_feed_follow
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT u.name userName, f.name feedName
FROM feeds_follows ff
JOIN users u ON u.id = ff.user_id
JOIN feeds f ON f.id = ff.feed_id
WHERE u.name = $1;