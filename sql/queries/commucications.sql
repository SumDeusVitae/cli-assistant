-- name: CreateComun :exec
INSERT INTO communications (id, created_at, updated_at, model, question, reply, user_id)
VALUES (?, ?, ?, ?, ?, ?, ?);
--

-- name: GetComunsById :one
SELECT * FROM communications WHERE id = ?;
--

-- name: GetComunsByUser :many
SELECT * FROM communications WHERE user_id = ?;

-- name: UpdateReply :exec
UPDATE communications
SET updated_at = ?,
    reply = ?
WHERE id = ?;