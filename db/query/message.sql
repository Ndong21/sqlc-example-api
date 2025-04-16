-- name: CreateThread :one
INSERT INTO "thread" (topic)
VALUEs ($1)
RETURNING *;

-- name: CreateMessage :one
INSERT INTO message (thread_id, sender, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateMessage :one
UPDATE message
SET content = $1
WHERE id = $2
RETURNING *;

-- name: GetMessageByID :one
SELECT * FROM message
WHERE id = $1;

-- name: GetMessagesByThread :many
SELECT * FROM message
WHERE thread_id = $1
ORDER BY created_at DESC;

-- name: DeleteMessageByID :exec
DELETE FROM message
WHERE id = $1;