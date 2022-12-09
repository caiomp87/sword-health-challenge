-- name: CreateTask :exec
INSERT INTO tasks (id, name, summary, performed, createdAt, performedAt, user_id) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateTask :exec
UPDATE tasks SET name = ?, summary = ? WHERE id = ?;

-- name: DoneTask :exec
UPDATE tasks SET performed = ?, performedAt = ? WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;

-- name: FindTaskById :one
SELECT * FROM tasks WHERE id = ?;

-- name: FindAllTasks :many
SELECT * FROM tasks;

-- name: CreateUser :exec
INSERT INTO users (id, name, type, email, passwordHash, createdAt) VALUES (?, ?, ?, ?, ?, ?);

-- name: FindUserById :one
SELECT * FROM users WHERE id = ?;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = ?;
