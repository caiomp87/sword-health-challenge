-- name: CreateTask :exec
INSERT INTO tasks (id, name, summary, performed, createdAt, performedAt, user_id) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateTask :exec
UPDATE tasks SET name = ?, summary = ? WHERE id = ? AND user_id = ?;

-- name: DoneTask :exec
UPDATE tasks SET performed = ?, performedAt = ? WHERE id = ? AND user_id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;

-- name: FindTaskByIDAndUserID :one
SELECT * FROM tasks WHERE id = ? AND user_id = ?;

-- name: FindTaskByID :one
SELECT * FROM tasks WHERE id = ?;

-- name: FindAllTasks :many
SELECT * FROM tasks;

-- name: FindAllTasksByUserID :many
SELECT * FROM tasks WHERE user_id = ?;

-- name: CreateUser :exec
INSERT INTO users (id, name, type, email, passwordHash, createdAt) VALUES (?, ?, ?, ?, ?, ?);

-- name: FindUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = ?;
