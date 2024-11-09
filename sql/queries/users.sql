-- name: CreateUser :exec
INSERT INTO users (id, created_at, updated_at, login, email, hashed_password, api_key)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);
--

-- name: GetUser :one
SELECT * FROM users WHERE api_key = ?;
--

-- name: DeleteUsers :exec
DELETE FROM users;
--

-- name: GetUserByLogin :one
SELECT * FROM users where login = ?;
--
