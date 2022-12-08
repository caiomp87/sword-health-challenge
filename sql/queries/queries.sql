-- name: CreateTask :exec
INSERT INTO tasks (id, name, summary) VALUES (?, ?, ?);

-- name: UpdateTask :exec
UPDATE tasks SET name = ?, summary = ?, performedAt = ? WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;

-- name: GetTaskById :one
SELECT * FROM tasks WHERE id = ?;

-- name: GetTasks :many
SELECT * FROM tasks;
