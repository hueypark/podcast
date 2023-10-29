-- name: GetFeed :one
SELECT title,
	description,
	updated
FROM feeds
WHERE title = ?
LIMIT 1;

-- name: ListFeeds :many
SELECT title,
	description,
	updated
FROM feeds;

-- name: CreateFeed :one
INSERT INTO feeds (title, description, link, feed_link, updated)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE title = ?;