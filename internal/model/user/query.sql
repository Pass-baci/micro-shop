-- name: GetUser :one
SELECT id, username, password, create_time, update_time  FROM user
WHERE id = ? LIMIT 1;

-- name: ListUser :many
SELECT id, username, password, create_time, update_time FROM user
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO user (
    id, username, password
) VALUES (?, ?, ?);

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;